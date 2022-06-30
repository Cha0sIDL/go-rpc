package main

import (
	"bytes"
	"context"
	"flag"
	"github.com/smallnest/rpcx/server"
	"io"
	"net/http"
)

type Args struct {
	Method string
	Url    string
	Body   []byte
}

type Http struct {
}

func (t *Http) Send(ctx context.Context, args *Args, Reply *[]byte) error {
	req, err := http.NewRequest(args.Method, args.Url, bytes.NewReader(args.Body))
	if err != nil {
		panic("request error")
	}
	client := &http.Client{}
	req.Header.Set("User-Agent", "User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)")
	response, err := client.Do(req)
	*Reply, _ = io.ReadAll(response.Body)
	return nil
}

var (
	addr = flag.String("addr", ":8972", "server address")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	s.RegisterName("Http", new(Http), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
