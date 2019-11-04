package main

import (
	"log"
	"math/rand"

	service "echo_v1"
)

func init() {
	var err error

	pattern := new(template1)
	pattern.setup = newServeContext
	pattern.teardown = disposeServeContext
	if pattern.proc, err = getServeFunc(); err != nil {
		log.Fatalln(err)
	}
	gservice = pattern
}

func getServeFunc() (fnServe, error) {
	return echoWithServeContext, nil
}

func newServeContext() (interface{}, error) {
	return uintptr(rand.Uint32()), nil
}

func disposeServeContext(interface{}) error {
	return nil
}

func echoWithServeContext(s interface{}, req *service.Request) (res *service.Response, ready bool, err error) {
	ready = true
	res = new(service.Response)
	if req != nil {
		res.Chunk = req.GetChunk()
	}
	return
}
