package mapepirego

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	godotenv.Load()

	var err error

	job, connectResponse, err := NewJob(
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("pass"),
		WithInsecureSkipVerify(),
	)
	defer job.Close()
	if err != nil {
		t.Error(err)
	}

	if connectResponse.Success == false {
		t.Error("connection unsuccessfull")
	}
}

func TestSelect(t *testing.T) {
	godotenv.Load()

	var err error

	job, connectResponse, err := NewJob(
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("pass"),
		WithInsecureSkipVerify(),
	)
	defer job.Close()
	if err != nil {
		t.Error(err)
	}
	if connectResponse.Success == false {
		t.Error("failed to connect to server")
	}

	testSchema := os.Getenv("test_schema")
	testTable := os.Getenv("test_table")

	sqlString := fmt.Sprintf("create table %s.%s ( a int, b char(10), c varchar(64))", testSchema, testTable)
	query := job.Query(sqlString)
	var createResult Result
	err = query.Execute(&createResult)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("create: %+v\n\n", createResult)

	var selectResult ResultWithData[struct {
		A float64
		B string
		C string
	}]
	sqlString = fmt.Sprintf("select * from %s.%s", testSchema, testTable)
	query = job.Query(sqlString)
	err = query.ExecuteSelect(&selectResult)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("select: %+v\n\n", selectResult)

	sqlString = fmt.Sprintf("insert into %s.%s values (5, '5', '6'), (6, '7', '8')", testSchema, testTable)
	query = job.Query(sqlString)
	var insertResult Result
	err = query.Execute(&insertResult)
	if err != nil {
		t.Error(err)
	}
	if insertResult.UpdateCount != 2 {
		t.Error("insert didn't update 2 rows")
	}
	fmt.Printf("insert: %+v\n\n", insertResult)

	sqlString = fmt.Sprintf("delete from %s.%s where a = 6", testSchema, testTable)
	query = job.Query(sqlString)
	var deleteResult Result
	err = query.Execute(&deleteResult)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("delete: %+v\n\n", deleteResult)
}
