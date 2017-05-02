// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log4go

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
)

// This log writer sends output to a socket
type SocketLogWriter struct {
	rec chan *LogRecord
	sync.Once
}

// This is the SocketLogWriter's output method
func (w *SocketLogWriter) LogWrite(rec *LogRecord) {
	defer func() {
		if e := recover(); e != nil {
			js, err := json.Marshal(rec)
			if err != nil {
				fmt.Printf("json error: %s", err)
				return
			}
			fmt.Printf("log channel has been closed. " + string(js) + "\n")
		}
	}()

	w.rec <- rec
}

func (w *SocketLogWriter) Close() {
	w.Once.Do(func() {
		close(w.rec)
	})
}

func NewSocketLogWriter(proto, hostport string) *SocketLogWriter {
	var w = &SocketLogWriter{}
	sock, err := net.Dial(proto, hostport)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewSocketLogWriter(%q): %w\n", hostport, err)
		return nil
	}

	w.rec = make(chan *LogRecord, LogBufferLength)

	go func() {
		defer func() {
			if sock != nil && proto == "tcp" {
				sock.Close()
			}
		}()

		for rec := range w.rec {
			// Marshall into JSON
			js, err := json.Marshal(rec)
			if err != nil {
				fmt.Fprint(os.Stderr, "SocketLogWriter(%q): %w", hostport, err)
				return
			}

			_, err = sock.Write(js)
			if err != nil {
				fmt.Fprint(os.Stderr, "SocketLogWriter(%q): %w", hostport, err)
				return
			}
		}
	}()

	return w
}
