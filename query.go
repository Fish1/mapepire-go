package mapepirego

import (
	"encoding/json"
)

type Query struct {
	job *Job
	sql string
}

func (query *Query) ExecuteSelect(into any) error {
	data, err := query.execute()

	err = json.Unmarshal(data, into)
	if err != nil {
		return err
	}

	return nil
}

func (query *Query) ExecuteInsert() (*InsertResult, error) {
	data, err := query.execute()

	var result InsertResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (query *Query) ExecuteCreate() (*CreateResult, error) {
	data, err := query.execute()

	var result CreateResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
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
