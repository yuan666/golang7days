module Day1/cmd_test

go 1.17

require github.com/mattn/go-sqlite3 v1.14.13

require mylog v1.0.0-1 // indirect

require session v1.0.0-1 // indirect

require geeorm v1.0.0-1 // indirect

replace geeorm => ../geeorm/

replace session => ../session/

replace mylog => ../mylog/
