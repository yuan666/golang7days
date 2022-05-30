module Day1/geeorm

go 1.17

require mylog v1.0.0-1 // indirect

require session v1.0.0-1 // indirect

replace session => ../session/

replace mylog => ../mylog/
