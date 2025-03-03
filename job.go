package mapepirego

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

type ConnectionOptions struct {
	Host               string
	Port               int
	User               string
	Pass               string
	InsecureSkipVerify bool
}

type Job struct {
	connection    *websocket.Conn
	timeout       context.CancelFunc
	id            string
	next_query_id int64
}

func CreateJob(options ConnectionOptions) (*Job, ConnectResponse, error) {
	job := Job{}
	connectResponse, err := job.Connect(options)
	if err != nil {
		return nil, connectResponse, err
	}
	return &job, connectResponse, nil
}

func (job *Job) getNextQueryID() int64 {
	job.next_query_id += 1
	return job.next_query_id
}

func (job *Job) send(buffer []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	err := job.connection.Write(ctx, websocket.MessageText, buffer)
	if err != nil {
		return err
	}
	return nil
}

func (job *Job) receive() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	_, data, err := job.connection.Read(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (job *Job) Query(sql string) Query {
	return Query{
		job: job,
		sql: sql,
	}
}

func (job *Job) Connect(options ConnectionOptions) (ConnectResponse, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: options.InsecureSkipVerify,
			},
		},
	}

	var ctx context.Context
	ctx, job.timeout = context.WithTimeout(context.Background(), time.Minute)

	authBuffer := []byte(fmt.Sprintf("%s:%s", options.User, options.Pass))
	auth := base64.StdEncoding.EncodeToString(authBuffer)

	header := http.Header{}
	header.Set("authorization", "Basic "+auth)

	dialOptions := &websocket.DialOptions{
		HTTPHeader: header,
		HTTPClient: httpClient,
	}

	url := fmt.Sprintf("wss://%s:%d/db/", options.Host, options.Port)

	var err error
	job.connection, _, err = websocket.Dial(ctx, url, dialOptions)
	if err != nil {
		return ConnectResponse{}, err
	}

	connectRequest, err := createConnectRequest(job.getNextQueryID())
	if err != nil {
		return ConnectResponse{}, err
	}

	err = job.send(connectRequest)
	if err != nil {
		return ConnectResponse{}, err
	}

	data, err := job.receive()
	if err != nil {
		return ConnectResponse{}, err
	}

	var connectResponse ConnectResponse
	err = json.Unmarshal(data, &connectResponse)
	if err != nil {
		return ConnectResponse{}, err
	}

	if connectResponse.Success == false {
		return ConnectResponse{}, errors.New("connection response: success = false")
	}

	job.id = connectResponse.Id

	return connectResponse, nil
}

func (job *Job) Close() error {
	return job.connection.CloseNow()
}
