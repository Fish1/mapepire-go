package mapepirego

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
)

func GetCertificate(options *ConnectionOptions) ([]*x509.Certificate, error) {
	conf := &tls.Config{
		InsecureSkipVerify: options.InsecureSkipVerify,
	}
	url := fmt.Sprintf("%s:%d", options.Host, options.Port)
	tlsConnection, err := tls.Dial("tcp", url, conf)
	if err != nil {
		return nil, err
	}
	defer tlsConnection.Close()

	return tlsConnection.ConnectionState().PeerCertificates, nil
}
