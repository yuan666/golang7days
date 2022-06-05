module Day2

go 1.17

require github.com/mattn/go-sqlite3 v1.14.13
require session v1.0.0-1

require dialect v1.0.0-1

require mylog v1.0.0-1

require schema v1.0.0-1 // indirect

replace schema => ./schema/

replace mylog => ./mylog/

replace dialect => ./dialect/

replace session => ./session/
