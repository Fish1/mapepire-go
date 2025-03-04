package mapepirego

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/coder/websocket"
)

type JobOption func(*Job)

func WithInsecureSkipVerify() JobOption {
	return func(job *Job) {
		job.insecureSkipVerify = true
	}
}

func WithFetchCertificate() JobOption {
	return func(job *Job) {
		job.fetchCertificate()
	}
}

type Job struct {
	connection         *websocket.Conn
	timeout            context.CancelFunc
	id                 string
	next_query_id      int64
	host               string
	port               string
	user               string
	pass               string
	certificate        string
	insecureSkipVerify bool
}

func NewJob(host string, port string, user string, pass string, opts ...JobOption) (*Job, ConnectResponse, error) {
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, ConnectResponse{}, errors.New("port must be a valid number")
	}

	job := Job{
		host: host,
		port: port,
		user: user,
		pass: pass,
	}

	for _, opt := range opts {
		opt(&job)
	}

	connectResponse, err := job.Connect()
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

func (job *Job) Connect() (ConnectResponse, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: job.insecureSkipVerify,
			},
		},
	}

	var ctx context.Context
	ctx, job.timeout = context.WithTimeout(context.Background(), time.Minute)

	authBuffer := []byte(fmt.Sprintf("%s:%s", job.user, job.pass))
	auth := base64.StdEncoding.EncodeToString(authBuffer)

	header := http.Header{}
	header.Set("authorization", "Basic "+auth)

	dialOptions := &websocket.DialOptions{
		HTTPHeader: header,
		HTTPClient: httpClient,
	}

	url := fmt.Sprintf("wss://%s:%s/db/", job.host, job.port)

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

func (job *Job) fetchCertificate() error {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	url := fmt.Sprintf("%s:%s", job.host, job.port)
	tlsConnection, err := tls.Dial("tcp", url, conf)
	if err != nil {
		return err
	}
	defer tlsConnection.Close()

	job.certificate = string(tlsConnection.ConnectionState().PeerCertificates[0].Raw)
	return nil
}
