package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"thrift/example"

	"github.com/apache/thrift/lib/go/thrift"
)

// FormatData 是实现了rpc服务的处理器processor
type FormatData struct {
}

// DoFormat 实现了thrift文件中定义的服务
func (s *FormatData) DoFormat(ctx context.Context, data *example.Data) (r *example.Data, err error) {
	var rData example.Data

	// 转化为go struct自动大写
	rData.Text = strings.ToUpper(data.Text)
	return &rData, nil
}

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {
	// 定义一个处理器，赋值给实现了服务的接口
	handler := &FormatData{}
	processor := example.NewFormatDataProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocalFactory := thrift.NewTBinaryProtocolFactoryDefault()
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocalFactory)
	fmt.Println("Running at:", HOST+":"+PORT)
	server.Serve()
}
