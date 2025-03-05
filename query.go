package mapepirego

import (
	"encoding/json"
)

type Query struct {
	job           *Job
	sql           string
	int_params    []int
	string_params []string
}

func (query *Query) Execute(into any) error {
	data, err := query.execute()
	if err != nil {
		return err
	}

	if into == nil {
		return nil
	}

	err = json.Unmarshal(data, &into)
	if err != nil {
		return err
	}

	return nil
}

func (query *Query) execute() ([]byte, error) {
	queryRequest, err := createSqlRequest(query.job.getNextQueryID(), query.sql)
	if err != nil {
		return nil, err
	}

	err = query.job.send(queryRequest)
	if err != nil {
		return nil, err
	}

	data, err := query.job.receive()
	if err != nil {
		return nil, err
	}

	return data, nil
}
