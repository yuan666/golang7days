package session

import (
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTabble(t *testing.T) {
	//s := NewSession().Model(&User{})
	/*
		s:=New(engine.db, engine.dialect)

		_ = s.DropTable()
		_ = s.CreateTable()
		if !s.HasTable() {
			t.Fatal("Failed to create table")
		}
	*/

}
