package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    //go语言的类型=>数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{}) //返回某个表是否存在的sql语句
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
