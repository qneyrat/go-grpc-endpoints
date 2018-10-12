package gen

import (
	internalgen "github.com/qneyrat/go-grpc-endpoints/internal/gen"
	googleproto "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	gen "github.com/golang/protobuf/protoc-gen-go/generator"
	"github.com/golang/protobuf/proto"
	internaltmpl "github.com/qneyrat/go-grpc-endpoints/internal/tmpl"
	"github.com/qneyrat/go-grpc-endpoints/protoc-gen-goendpoints/tmpl"
	"go/format"
)

type generator struct {
	*internalgen.Generator
}

func New() *generator {
	return &generator{Generator: internalgen.New()}
}

type Service struct {
	Name string
	Methods []*googleproto.MethodDescriptorProto
}

type Data struct {
	Package string
	Services []Service
}

func (g *generator) Make(protoFile *googleproto.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	data := Data{
		Package: protoFile.GetPackage(),
		Services: make([]Service, len(protoFile.Service)),
	}

	for _, service := range protoFile.Service {
		data.Services = append(data.Services, Service{
			Name: gen.CamelCase(service.GetName()),
			Methods: service.Method,
		})
	}

	b, err := tmpl.Asset("endpoints.pb.go.tmpl")
	if err != nil {
		return &plugin.CodeGeneratorResponse_File{}, err
	}

	t, err := internaltmpl.NewTempale(b)
	if err != nil {
		return &plugin.CodeGeneratorResponse_File{}, err
	}

	buf, err := t.Eval(data)
	if err != nil {
		return &plugin.CodeGeneratorResponse_File{}, err
	}

	g.P(buf.String())

	formatted, err := format.Source(g.Bytes())
	if err != nil {
		return &plugin.CodeGeneratorResponse_File{}, err
	}

	file := &plugin.CodeGeneratorResponse_File{
		Name:    proto.String(g.ProtoFileBaseName(*protoFile.Name) + ".endpoints.pb.go"),
		Content: proto.String(string(formatted)),
	}
	return file, nil
}

func (g *generator) Generate() {
	g.Generator.Generate(g)
}
