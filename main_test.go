package mapepirego

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	options := ConnectionOptions{
		InsecureSkipVerify: true,
	}

	err := godotenv.Load()
	if err != nil {
		t.Error(err)
	}

	host, _ := os.LookupEnv("host")
	port, _ := os.LookupEnv("port")
	user, _ := os.LookupEnv("user")
	pass, _ := os.LookupEnv("pass")

	options.Host = host
	options.Port, err = strconv.Atoi(port)
	if err != nil {
		t.Error(err)
	}
	options.User = user
	options.Pass = pass

	fmt.Println(options)

	job := Job{}

	connectResponse, err := job.Connect(options)
	defer job.Close()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(connectResponse)

	query := job.Query("select * from JENDERS1.MYTABLE")
	selectResult, err := query.ExecuteSelect()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(selectResult)

	query = job.Query("insert into JENDERS1.MYTABLE1 values (1, 2, 3, 4, 5)")
	insertResult, err := query.ExecuteInsert()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(insertResult)

	query = job.Query("select * from JENDERS1.MYTABLE1")
	selectResult, err = query.ExecuteSelect()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(selectResult)

	query = job.Query("create table JENDERS1.MYTABLE2 ( a int, b char(10), c varchar(64))")
	createResult, err := query.ExecuteCreate()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(createResult)

	query = job.Query("create table JENDERS1.MYTABLE3 ( a int, b char(10), c varchar(64))")
	createResult, err = query.ExecuteCreate()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(createResult)

	query = job.Query("insert into JENDERS1.MYTABLE2 values (1, 'aa', 'bbb')")
	insertResult, err = query.ExecuteInsert()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(insertResult)

	query = job.Query("select * from JENDERS1.MYTABLE2")
	selectResult, err = query.ExecuteSelect()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(selectResult)
}
