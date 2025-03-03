package mapepirego

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	godotenv.Load()

	var err error

	options := ConnectionOptions{
		InsecureSkipVerify: true,
	}

	options.Host, _ = os.LookupEnv("host")
	options.User, _ = os.LookupEnv("user")
	options.Pass, _ = os.LookupEnv("pass")

	port, _ := os.LookupEnv("port")
	options.Port, err = strconv.Atoi(port)
	if err != nil {
		t.Error(err)
	}

	job := Job{}
	connectResponse, err := job.Connect(options)
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

	options := ConnectionOptions{
		InsecureSkipVerify: true,
	}

	options.Host, _ = os.LookupEnv("host")
	options.User, _ = os.LookupEnv("user")
	options.Pass, _ = os.LookupEnv("pass")

	port, _ := os.LookupEnv("port")
	options.Port, err = strconv.Atoi(port)
	if err != nil {
		t.Error(err)
	}

	job, connectResponse, err := CreateJob(options)
	defer job.Close()
	if err != nil || connectResponse.Success == false {
		t.Error(err)
	}

	query := job.Query("create table JENDERS1.MYTABLE2 ( a int, b char(10), c varchar(64))")
	_, err = query.ExecuteCreate()
	if err != nil {
		t.Error(err)
	}

	var selectResult SelectResult[struct {
		A float64
		B string
		C string
	}]
	query = job.Query("select * from JENDERS1.MYTABLE2")
	err = query.ExecuteSelect(&selectResult)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(selectResult)
}
