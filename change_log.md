# log4go #
---
*golang logger of log4j*

## dev list ##
---

- 2017/04/28
   > bug fix
   * add recover logic in file write function
   * use maxbackup instead of 999

- 2017/04/24
   > improvement
   * add filename in log4go.go:intLogf & intLogc
   

- 2017/03/07
	> bug fix: log4go panic when cannot open new log file
	>
	> version: 3.0.2
	
- 2017/02/09
	> bug fix: just closed once in log4go.go
	>
	> bug fix: add select-case in every LogWrite to avoid panic when log channel is closed.


- 2016/09/21
	> add Logger{FilterMap, minLevel}
	>
	> add l4g.Close for examples/XMLConfigurationExample.go
	>
	> delete redundant l4g.Close for examples/FileLogWriter_Manual.go
	>
	> modify the return value to nil of log.Warn & log.Error & log.Critical
	>
	> add Chinese remark for some key functions

