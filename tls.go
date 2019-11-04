package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	tlsconfig *tls.Config    = nil
	autority  *x509.CertPool = nil
)

func cacheTLSFromEnv() error {
	keypath, defined := os.LookupEnv("KEY_PATH")
	if defined == false {
		return errors.New("env KEY_PATH undefined")
	}
	certpath, defined := os.LookupEnv("CERT_PATH")
	if defined == false {
		return errors.New("env CERT_PATH undefined")
	}
	return cacheTLSFromParams(keypath, certpath)
}

func cacheTLSFromParams(keypath, certpath string) error {
	if tlsconfig != nil {
		return nil
	}

	cert, err := tls.LoadX509KeyPair(certpath, keypath)
	if err != nil {
		return err
	}
	tlsconfig = &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{cert},
		// ClientAuth: tls.RequireAndVerifyClientCert,
		// CipherSuites: []uint16{
		// 	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		// 	tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		// 	tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		// 	tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
		// },
	}
	// tlsconfig.RootCAs, err = makeCertPool(certpath)
	return err
}

func makeCertPool(certpath string) (*x509.CertPool, error) {
	if autority != nil {
		return autority, nil
	}
	autority = x509.NewCertPool()
	blob, err := ioutil.ReadFile(certpath)
	if err != nil {
		return nil, err
	}
	if autority.AppendCertsFromPEM(blob) == false {
		return nil, fmt.Errorf("failed to add cert from pem: %s", certpath)
	}
	return autority, nil
}
