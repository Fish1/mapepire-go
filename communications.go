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

type SelectResultColumnMetaData struct {
	DisplaySize int    `json:"display_size"`
	Label       string `json:"label"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Precision   int    `json:"precision"`
	Scale       int    `json:"scale"`
}

type SelectResultParameterDetail struct {
	Type      string `json:"type"`
	Mode      string `json:"mode"`
	Precision int    `json:"precision"`
	Scale     int    `json:"scale"`
	Name      string `json:"name"`
}

type SelectResultMetaData struct {
	ColumnCount int                           `json:"column_count"`
	Columns     []SelectResultColumnMetaData  `json:"columns"`
	Parameters  []SelectResultParameterDetail `json:"parameters"`
	Job         string                        `json:"job"`
}

type SelectResultOutputParms struct {
}

type SelectResult[data any] struct {
	Id             string                  `json:"id"`
	MetaData       SelectResultMetaData    `json:"metadata"`
	IsDone         bool                    `json:"is_done"`
	HasResults     bool                    `json:"has_results"`
	UpdateCount    int                     `json:"update_count"`
	Data           []data                  `json:"data"`
	ParameterCount int                     `json:"parameter_count"`
	OutputParams   SelectResultOutputParms `json:"output_parms"`
}

type InsertResult struct {
	Id            string `json:"id"`
	ExecutionTime int    `json:"execution_time"`
	HasResults    bool   `json:"has_results"`
	Success       bool   `json:"success"`
	UpdateCount   int    `json:"update_count"`
}

type CreateResult struct {
	Id            string `json:"id"`
	Success       bool   `json:"success"`
	UpdateCount   int    `json:"update_count"`
	Error         string `json:"error"`
	ExecutionTime int    `json:"execution_time"`
	HasResults    bool   `json:"has_results"`
}
