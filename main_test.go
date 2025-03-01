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

	err = job.Connect(options)
	defer job.Close()
	if err != nil {
		t.Error(err)
	}

	query := job.Query("select * from JENDERS1.MYTABLE")
	err = query.Execute()
	if err != nil {
		t.Error(err)
	}
}
