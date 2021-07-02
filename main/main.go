package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	"youngrpc"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	youngrpc.Accept(l)
}

func main() {
	//addr := make(chan string)
	//go startServer(addr)
	//
	//// 模拟一个简单的客户端
	//conn, _ := net.Dial("tcp", <-addr)
	//defer func() { _ = conn.Close() }()
	//
	//time.Sleep(time.Second)
	//// 发送Option进行协议交互
	//_ = json.NewEncoder(conn).Encode(youngrpc.DefaultOption)
	//cc := codec.NewGobCodec(conn)
	//
	//// 发送请求 & 接收响应
	//for i := 0; i < 5; i++ {
	//	h := &codec.Header{
	//		ServiceMethod: "Foo.Sum",
	//		Seq:           uint64(i),
	//	}
	//	_ = cc.Write(h, fmt.Sprintf("youngrpc req %d", h.Seq))    // 构造请求
	//
	//	_ = cc.ReadHeader(h)    // 处理请求
	//	log.Println("head:", h)
	//	var reply string
	//	//_ = cc.ReadHeader(h)
	//	_ = cc.ReadBody(&reply)
	//	log.Println("reply:", reply)
	//}
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)
	client, _ := youngrpc.Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// 发起请求并接收响应
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("youngrpc reqs %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
