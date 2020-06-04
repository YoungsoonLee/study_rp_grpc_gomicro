package main

import (
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/labstack/gommon/log"
)

type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeserver := new(TimeServer)
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	http.Serve(l, nil)
}
