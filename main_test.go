package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var conn *grpc.ClientConn

func grpcConnT(t *testing.T) *grpc.ClientConn {
	if conn != nil {
		return conn
	}

	creds := credentials.NewClientTLSFromCert(autority, "localhost")
	conn, err := grpc.Dial("localhost:8443", grpc.WithTransportCredentials(creds))
	if err != nil {
		t.Fatal(err)
	}
	return conn
}

func TestGRPCDial(t *testing.T) {
	conn := grpcConnT(t)
	if conn == nil {
		t.FailNow()
	}
	t.Log(conn.Target())
}

func TestMain(m *testing.M) {
	log.SetFlags(log.Lshortfile)

	// run with TLS
	workspace, err := os.Getwd()
	log.Println(workspace)

	hostname := "localhost"
	keypath := filepath.Join(workspace, "scp227.key")
	certpath := filepath.Join(workspace, "localhost.crt")

	if err = os.Setenv("KEY_PATH", keypath); err != nil {
		log.Fatalln(err)
	}
	if err = os.Setenv("CERT_PATH", certpath); err != nil {
		log.Fatalln(err)
	}
	if err = cacheTLSFromEnv(); err != nil {
		log.Fatalln(err)
	}
	// make self-signed to test locally
	tlsconfig.ServerName = hostname
	if tlsconfig.RootCAs, err = makeCertPool(certpath); err != nil {
		log.Fatalln(err)
	}

	go serveH(":8443")
	time.Sleep(1 * time.Second)

	// close the connection
	defer func() {
		if conn == nil {
			return
		}
		conn.Close()
	}()
	os.Exit(m.Run())
}
