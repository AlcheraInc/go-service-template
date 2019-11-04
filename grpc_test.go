package main

import (
	"bytes"
	"context"
	"io"
	"testing"

	service "echo_v1"

	"google.golang.org/grpc"
)

func clientT1(t *testing.T, conn *grpc.ClientConn, req *service.Request) *service.Response {
	cli := service.NewEchoClient(conn)
	res, err := cli.Serve1(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func clientT2(t *testing.T, conn *grpc.ClientConn, requests <-chan *service.Request) (res *service.Response) {
	cli := service.NewEchoClient(conn)

	ostream, err := cli.Serve2(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	for req := range requests {
		if err = ostream.Send(req); err != nil {
			t.Fatal(err)
		}
	}
	if err = ostream.CloseSend(); err != nil {
		t.Fatal(err)
	}

	if res, err = ostream.CloseAndRecv(); err != nil {
		t.Fatal(err)
	}
	return res
}

func clientT3(t *testing.T, conn *grpc.ClientConn, req *service.Request) <-chan *service.Response {
	cli := service.NewEchoClient(conn)
	responses := make(chan *service.Response, 20)
	defer close(responses)

	istream, err := cli.Serve3(context.Background(), req)
	if err != nil {
		return responses
	}
	if err = istream.CloseSend(); err != nil {
		t.Fatal(err)
	}

	for {
		res, err := istream.Recv()
		if err == io.EOF {
			return responses
		}
		if err != nil {
			t.Fatal(err)
		}
		responses <- res
	}
}

func clientT4(t *testing.T, conn *grpc.ClientConn, requests <-chan *service.Request) <-chan *service.Response {
	cli := service.NewEchoClient(conn)
	responses := make(chan *service.Response, 20)
	defer close(responses)

	iostream, err := cli.Serve4(context.Background())
	if err != nil {
		return responses
	}

	for req := range requests {
		if err = iostream.Send(req); err != nil {
			t.Fatal(err)
		}
	}
	if err = iostream.CloseSend(); err != nil {
		t.Fatal(err)
	}

	for {
		res, err := iostream.Recv()
		if err == io.EOF {
			return responses
		}
		if err != nil {
			t.Fatal(err)
		}
		responses <- res
	}
}

func TestGRPCServe1Empty(t *testing.T) {
	req := new(service.Request)
	req.Chunk = nil

	res := clientT1(t, grpcConnT(t), req)
	if len(res.GetChunk()) != 0 {
		t.FailNow()
	}
}

func TestGRPCServe1String(t *testing.T) {
	req := new(service.Request)
	req.Chunk = []byte("hellworld")

	res := clientT1(t, grpcConnT(t), req)
	if chunk := res.GetChunk(); len(chunk) == 0 {
		t.FailNow()
	} else if bytes.Equal(chunk, req.Chunk) == false {
		t.FailNow()
	}
}

func TestGRPCServe2(t *testing.T) {
	requests := make(chan *service.Request, 5)
	for _, txt := range []string{"aaa", "bbb", "ccc", "ddd", "eee"} {
		req := new(service.Request)
		req.Chunk = []byte(txt)
		requests <- req
	}
	close(requests)

	res := clientT2(t, grpcConnT(t), requests)
	if chunk := res.GetChunk(); len(chunk) == 0 {
		t.FailNow()
	} else if bytes.Equal(chunk, []byte("aaa")) == false {
		t.FailNow()
	}
}

func TestGRPCServe3(t *testing.T) {
	req := new(service.Request)
	req.Chunk = []byte("hellworld")

	responses := clientT3(t, grpcConnT(t), req)
	if res := <-responses; string(res.GetChunk()) != "hellworld" {
		t.FailNow()
	}
}

func TestGRPCServe4(t *testing.T) {
	requests := make(chan *service.Request, 5)
	for _, txt := range []string{"aaa", "bbb", "ccc", "ddd", "eee"} {
		req := new(service.Request)
		req.Chunk = []byte(txt)
		requests <- req
	}
	close(requests)

	responses := clientT4(t, grpcConnT(t), requests)
	if res := <-responses; string(res.GetChunk()) != "aaa" {
		t.FailNow()
	}
	if res := <-responses; string(res.GetChunk()) != "bbb" {
		t.FailNow()
	}
	if res := <-responses; string(res.GetChunk()) != "ccc" {
		t.FailNow()
	}
	if res := <-responses; string(res.GetChunk()) != "ddd" {
		t.FailNow()
	}
	if res := <-responses; string(res.GetChunk()) != "eee" {
		t.FailNow()
	}
}
