package main

import (
	"log"
	"net"
	"os"
)

func init() {
	// just check some network features ...
	hostname, _ := os.Hostname()
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		log.Fatalln(err)
	}
	for i := range addrs {
		log.Println(hostname, addrs[i])
	}
}

func serveH(address string) {
	keypath, defined := os.LookupEnv("KEY_PATH")
	if defined == false {
		serve1(address)
		return
	}
	certpath, defined := os.LookupEnv("CERT_PATH")
	if defined == false {
		serve1(address)
		return
	}
	log.Println("cert file:", certpath)
	serve2(address, keypath, certpath)
}

func main() {
	if err := cacheTLSFromEnv(); err != nil {
		log.Fatalln(err)
	}

	workspace, _ := os.Getwd()
	log.Println("workspace:", workspace)

	serveH(":80")
}
