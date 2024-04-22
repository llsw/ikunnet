// Code generated by Kitex v0.9.1. DO NOT EDIT.

package transportservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/llsw/ikunet/kitex_gen/transport"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Call": kitex.NewMethodInfo(
		callHandler,
		newCallArgs,
		newCallResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	transportServiceServiceInfo                = NewServiceInfo()
	transportServiceServiceInfoForClient       = NewServiceInfoForClient()
	transportServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return transportServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return transportServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return transportServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "TransportService"
	handlerType := (*transport.TransportService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "transport",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func callHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(transport.Transport)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(transport.TransportService).Call(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CallArgs:
		success, err := handler.(transport.TransportService).Call(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CallResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCallArgs() interface{} {
	return &CallArgs{}
}

func newCallResult() interface{} {
	return &CallResult{}
}

type CallArgs struct {
	Req *transport.Transport
}

func (p *CallArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(transport.Transport)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CallArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CallArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CallArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CallArgs) Unmarshal(in []byte) error {
	msg := new(transport.Transport)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CallArgs_Req_DEFAULT *transport.Transport

func (p *CallArgs) GetReq() *transport.Transport {
	if !p.IsSetReq() {
		return CallArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CallArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CallArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CallResult struct {
	Success *transport.Transport
}

var CallResult_Success_DEFAULT *transport.Transport

func (p *CallResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(transport.Transport)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CallResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CallResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CallResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CallResult) Unmarshal(in []byte) error {
	msg := new(transport.Transport)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CallResult) GetSuccess() *transport.Transport {
	if !p.IsSetSuccess() {
		return CallResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CallResult) SetSuccess(x interface{}) {
	p.Success = x.(*transport.Transport)
}

func (p *CallResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CallResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Call(ctx context.Context, Req *transport.Transport) (r *transport.Transport, err error) {
	var _args CallArgs
	_args.Req = Req
	var _result CallResult
	if err = p.c.Call(ctx, "Call", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
