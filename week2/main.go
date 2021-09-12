package main

import (
	"database/sql"
	"errors"
	"fmt"
	es "github.com/pkg/errors"
)

// 应该抛给上层处理

func Query(sqls string) error {
	return es.Wrap(sql.ErrNoRows, "ErrNoRows")
}

func main() {
	err := Query("select * from tale")
	if  err!=nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("type: %T \nerror: %+v", es.Cause(err), es.Cause(err))
			fmt.Printf("%+v", err)
		} else {
			fmt.Printf("error: %+v", err)
		}
		return
	}

	println("The Show Must Go On")
}