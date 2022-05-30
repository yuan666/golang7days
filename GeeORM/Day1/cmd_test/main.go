package main

import (
	"fmt"
	"geeorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "geeNew.db")

	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()

	result, _ := s.Raw("INSERT INTO User(Name) values (?),(?)", "Tom", "Jack").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
/***
database/sql的基础操作

测试结果：
cmd_test git:(master) ✗ go build main.go
➜  cmd_test git:(master) ✗ ./main
[info ] 2022/05/30 16:17:46 geeorm.go:28: Connect database success
[info ] 2022/05/30 16:17:46 raw.go:39: DROP TABLE IF EXISTS User;  []
[info ] 2022/05/30 16:17:46 raw.go:39: CREATE TABLE User(Name text);  []
[info ] 2022/05/30 16:17:46 raw.go:39: CREATE TABLE User(Name text);  []
[error] 2022/05/30 16:17:46 raw.go:42: table User already exists
[info ] 2022/05/30 16:17:46 raw.go:39: INSERT INTO User(Name) values (?),(?)  [Tom Jack]
Exec success, 2 affected
[info ] 2022/05/30 16:17:46 geeorm.go:36: Close database success
**/