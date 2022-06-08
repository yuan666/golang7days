package clause

//实现结构体拼接，各个独立的子句

type Clause struct {
	sql map[Type]string
	sqlVars map[Type]interface{}
}

type  Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

