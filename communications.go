package mapepirego

import "encoding/json"

type connectRequest struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Technique   string `json:"technique"`
	Application string `json:"applicatoin"`
	// Props       string `json:"props"`
}

func createConnectRequest() ([]byte, error) {
	return json.Marshal(connectRequest{
		Id:          "1",
		Type:        "connect",
		Technique:   "tcp",
		Application: "mapepire-go",
	})
}

type connectResponse struct {
	Id             string `json:"id"`
	Job            string `json:"job"`
	Success        bool   `json:"success"`
	Execution_Time int    `json:"execution_time"`
}

type queryRequest struct {
	Id         string        `json:"id"`
	Type       string        `json:"type"`
	Sql        string        `json:"sql"`
	Terse      bool          `json:"terse"`
	Rows       int           `json:"rows"`
	Parameters []interface{} `json:"parameters"`
}

func createQueryRequest(sql string) ([]byte, error) {
	return json.Marshal(queryRequest{
		Id:         "1",
		Type:       "sql",
		Sql:        sql,
		Terse:      false,
		Rows:       50,
		Parameters: []interface{}{},
	})
}

type queryResultColumnMetaData struct {
	DisplaySize int    `json:"display_size"`
	Label       string `json:"label"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Precision   int    `json:"precision"`
	Scale       int    `json:"scale"`
}

type queryResultParameterDetail struct {
	Type      string `json:"type"`
	Mode      string `json:"mode"`
	Precision int    `json:"precision"`
	Scale     int    `json:"scale"`
	Name      string `json:"name"`
}

type queryResultMetaData struct {
	ColumnCount int                          `json:"column_count"`
	Columns     []queryResultColumnMetaData  `json:"columns"`
	Parameters  []queryResultParameterDetail `json:"parameters"`
	Job         string                       `json:"job"`
}

type queryResultOutputParms struct {
}

type queryResult struct {
	MetaData       queryResultMetaData    `json:"metadata"`
	IsDone         bool                   `json:"is_done"`
	HasResults     bool                   `json:"has_results"`
	UpdateCount    int                    `json:"update_count"`
	Data           []interface{}          `json:"data"`
	ParameterCount int                    `json:"parameter_count"`
	OutputParams   queryResultOutputParms `json:"output_parms"`
}
