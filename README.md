# gotools
gotools directory store every go tools.
20151209 add logs module
logs rewrite base on the log in go source. 
logs can write logs to file or stdout.
logs file name's prefix can be set datetime or sequence. 
logs file can set the size of file. when the size of logs file exceed the provide size, the logs will use new file to write logs.
logs default write log to stdout. when user set the path of file, auto write log to file.
logs support INFO,DEBUG,WARING,ERROR,FATAL.when the level is FATAL, writing log will exit current run program.
1.for example(write to stdout)
logs.Debug("write to stdout")
2. write to file
logs.LogSetFilePath("log.txt")
logs.Debug("write to file")
