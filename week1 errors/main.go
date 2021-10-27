package main

import (
	sysErr "errors"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

var (
	ErrNoRows = sysErr.New("no records found")
)

type Student struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	stu, err := queryStudent("select * from students where id = ?", 0)
	if err != nil {
		log.Printf("failed to query students, %v", err.Error())
		return
	}
	log.Printf("student name is %s", stu.Name)
}

func queryStudent(sql string, id int64) (*Student, error) {
	if id == 1 {
		return &Student{}, nil
	}
	return nil, errors.Wrap(ErrNoRows, fmt.Sprintf("sql %s query no record", sql))
}
