package main

import (
	"fmt"
	"log"
	"net"
	"thrift/example"

	"github.com/apache/thrift/lib/go/thrift"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocketError", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, _ := transportFactory.GetTransport(tSocket)
	protocalFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := example.NewFormatDataClientFactory(transport, protocalFactory)
	if err = transport.Open(); err != nil {
		log.Fatalln("Error opening", HOST+":"+PORT)
	}
	defer transport.Close()
	data := example.Data{Text: "hello, world"}
	d, err := client.DoFormat(nil, &data)
	fmt.Println(d.Text)
}
