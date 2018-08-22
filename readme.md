# go-grpc-endpoints
protoc plugin to auto-generate endpoints layer

Helpfull plugin to use composition on gRPC server.

## Installing and using

Install with `go get`
```
> $ go get github.com/qneyrat/go-grpc-endpoints/protoc-gen-goendpoints
> $ go install github.com/qneyrat/go-grpc-endpoints/protoc-gen-goendpoints
```

Add `--goendpoints_out=` flag to generage endpoints layer with protoc
```
> $ protoc -I . ./translator.proto --go_out=plugins=grpc:. --goendpoints_out=.:.   
```

## Plugin and endpoints layer usage

Plugin auto-generate strucs and func for endpoints layer from your proto files

EndpointsWrapper is gRPC server. It's implement gRPC server interface and call endpoints

Endpoints struct is endpoints group to crate EndpointsWrapper

You can just write EndpointFunc constructor to hydrate Endpoints struct:

> Create EndpointFunc constructor
```go 
func NewTranslateEndpoint(t translate.Translator) TranslateEndpointFunc {
  return func(ctx context.Context, req *proto.TranslateRequest) (*proto.TranslateResponse, error) {
     // use t.Translate here
  }
}
```

> Register EndpointsWrapper on gRPC server
```go
	translator := translate.NewGoogleTranslator()
	wrapper := NewTranslatorEndpointsWrapper(TranslatorEndpoints{
		TranslateEndpoint: NewTranslateEndpoint(translator),
	})

	s := grpc.NewServer()
	proto.RegisterTranslatorServer(s, wrapper)
	s.Serve(lis)
```

## Theory

ServiceX represents application services like database repository or domain service.

### gRPC Server without composition
```Go
type Server struct{
  service1 Service1
  service2 Service2
  service3 Service3
  // ...
}

func (s *Service) Method(ctx context.Context, req *proto.Req) (*proto.Res, error) {
  // use services here like s.service1.Foo()
}
```

### gRPC Server with composition and introduce endpoints layer
```Go
type EndpointFunc1 func(ctx context.Context, req *proto.Req) (*proto.Res, error)

func NewEndpointFunc1(service1 Service1,  service2 Service2) EndpointFunc1 {
  return func(ctx context.Context, req *proto.Req) (*proto.Res, error) {
      // use services here like s.service1.Foo()
  }
}

type Server struct{
  endpoint1 EndpointFunc1
  // ...
}

func (s *Service) Method(ctx context.Context, req *proto.Req) (*proto.Res, error) {
  return s.endpoint1(ctx, req)
}
```
