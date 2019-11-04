package main

import (
	"context"
	"io"

	service "echo_v1"
)

//
//	contextful service functions
//

type fnSetup func() (session interface{}, err error)
type fnTeardown func(session interface{}) error
type fnServe func(session interface{}, req *service.Request) (res *service.Response, ready bool, err error)

//
//	Basic ECHO implementation
//

type template1 struct {
	setup    fnSetup
	teardown fnTeardown
	proc     fnServe
}

func (impl *template1) Serve1(ctx context.Context, req *service.Request) (*service.Response, error) {
	s, err := impl.setup()
	if err != nil {
		return nil, err
	}
	defer impl.teardown(s)

	res, _, err := impl.proc(s, req)
	return res, err
}

func (impl *template1) Serve2(istream service.Echo_Serve2Server) error {
	s, err := impl.setup()
	if err != nil {
		return err
	}
	defer impl.teardown(s)

	for {
		req, err := istream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		res, _, err := impl.proc(s, req)
		if err != nil {
			return err
		}
		return istream.SendAndClose(res)
	}
}

func (impl *template1) Serve3(req *service.Request, ostream service.Echo_Serve3Server) error {
	s, err := impl.setup()
	if err != nil {
		return err
	}
	defer impl.teardown(s)

	res, _, err := impl.proc(s, req)
	if err != nil {
		return err
	}
	return ostream.Send(res)
}

func (impl *template1) Serve4(iostream service.Echo_Serve4Server) error {
	s, err := impl.setup()
	if err != nil {
		return err
	}
	defer impl.teardown(s)

	for {
		req, err := iostream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		res, ready, err := impl.proc(s, req)
		if err != nil {
			return err
		}
		if ready == false {
			continue
		}
		if err = iostream.Send(res); err != nil {
			return err
		}
	}
}
