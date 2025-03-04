package mapepirego

import (
	"encoding/json"
	"strconv"
)

type connectRequest struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Technique   string `json:"technique"`
	Application string `json:"applicatoin"`
	// Props       string `json:"props"`
}

func createConnectRequest(id int64) ([]byte, error) {
	return json.Marshal(connectRequest{
		Id:          strconv.FormatInt(id, 36),
		Type:        "connect",
		Technique:   "tcp",
		Application: "mapepire-go",
	})
}

type ConnectResponse struct {
	Id             string `json:"id"`
	Job            string `json:"job"`
	Success        bool   `json:"success"`
	Execution_Time int    `json:"execution_time"`
}

type sqlRequest struct {
	Id         string        `json:"id"`
	Type       string        `json:"type"`
	Sql        string        `json:"sql"`
	Terse      bool          `json:"terse"`
	Rows       int           `json:"rows"`
	Parameters []interface{} `json:"parameters"`
}

func createSqlRequest(id int64, sql string) ([]byte, error) {
	return json.Marshal(sqlRequest{
		Id:         strconv.FormatInt(id, 36),
		Type:       "sql",
		Sql:        sql,
		Terse:      false,
		Rows:       50,
		Parameters: []interface{}{},
	})
}

type Result struct {
	Id            string `json:"id"`
	Success       bool   `json:"success"`
	HasResults    bool   `json:"has_results"`
	UpdateCount   int    `json:"update_count"`
	ExecutionTime int    `json:"execution_time"`
	Error         string `json:"error"`
}

type ResultWithData[data any] struct {
	Id            string `json:"id"`
	Success       bool   `json:"success"`
	HasResults    bool   `json:"has_results"`
	UpdateCount   int    `json:"update_count"`
	ExecutionTime int    `json:"execution_time"`

	MetaData       ResultMetaData          `json:"metadata"`
	IsDone         bool                    `json:"is_done"`
	Data           []data                  `json:"data"`
	ParameterCount int                     `json:"parameter_count"`
	OutputParams   SelectResultOutputParms `json:"output_parms"`

	Error string `json:"error"`
}

type ResultColumnMetaData struct {
	DisplaySize int    `json:"display_size"`
	Label       string `json:"label"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Precision   int    `json:"precision"`
	Scale       int    `json:"scale"`
}

type ResultParameterDetail struct {
	Type      string `json:"type"`
	Mode      string `json:"mode"`
	Precision int    `json:"precision"`
	Scale     int    `json:"scale"`
	Name      string `json:"name"`
}

type ResultMetaData struct {
	ColumnCount int                     `json:"column_count"`
	Columns     []ResultColumnMetaData  `json:"columns"`
	Parameters  []ResultParameterDetail `json:"parameters"`
	Job         string                  `json:"job"`
}

type SelectResultOutputParms struct {
}
