package youngrpc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
	"youngrpc/codec"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int        // 魔术数字用于标识这是一个RPC请求
	CodecType   codec.Type // 客户端采用的编码方式
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

type Server struct{} // Server标识一个RPC Server

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer() // RPC Server的默认实例

// Accept方法接收监听的连接并处理请求，使用Server作为接收者
func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error", err)
			return
		}
		go server.ServerConn(conn)
	}
}

// 主调的通用Accept方法
func Accept(lis net.Listener) { DefaultServer.Accept(lis) }

// ServerConn处理连接，校验是否rpc请求
func (server *Server) ServerConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: options error: ", err)
		return
	}
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}
	server.serveCodec(f(conn))
}

var invalidRequest = struct{}{} // 当异常产生时作为占位符进行响应

// 读取请求、处理请求、回复请求
func (server *Server) serveCodec(cc codec.Codec) {
	sending := new(sync.Mutex) // 确保返回一个完整的响应
	wg := new(sync.WaitGroup)  // 等待所有请求完成
	for {
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break // 此时不可能修复，需要关闭连接
			}
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}
		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg)
	}
	wg.Wait()
	_ = cc.Close()
}

type request struct {
	h            *codec.Header // 请求头部
	argv, replyv reflect.Value // 请求参数
}

func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{h: h}

	req.argv = reflect.New(reflect.TypeOf("")) // 设定请求参数类型是string
	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}
	return req, nil
}

func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	sending.Lock()
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil{
		log.Println("rpc server: write response error:", err)
	}
}

func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup){
	// TODO: 应该调用服务发现方法获取正确的replyv
	defer wg.Done()
	log.Println("hhhh", req.h, req.argv.Elem())
	req.replyv = reflect.ValueOf(fmt.Sprintf("youngrpc resp %d", req.h.Seq))

	//req.h.Seq = req.h.Seq*10
	server.sendResponse(cc, req.h, req.replyv.Interface(), sending)
	//server.sendResponse(cc, req.h, "???", sending)
}
