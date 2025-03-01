package mapepirego

import (
	"encoding/json"
	"fmt"
)

type Query struct {
	job *Job
	sql string
}

func (query *Query) Execute() error {
	queryRequest, err := createQueryRequest(query.sql)
	if err != nil {
		return err
	}

	err = query.job.send(queryRequest)
	if err != nil {
		return err
	}

	data, err := query.job.receive()
	if err != nil {
		return err
	}

	var queryResponse interface{}
	err = json.Unmarshal(data, &queryResponse)
	if err != nil {
		return err
	}

	fmt.Println(queryResponse)

	return nil
}
