package gen

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/golang/protobuf/proto"
	googleproto "github.com/golang/protobuf/protoc-gen-go/descriptor"
	gen "github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

type Maker interface {
	Make(*googleproto.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)
}

type Generator struct {
	*gen.Generator
	indent string

	reader io.Reader
	writer io.Writer
}

func New() *Generator {
	return &Generator{
		Generator: gen.New(),
		reader:    os.Stdin,
		writer:    os.Stdout,
	}
}

func (g *Generator) error(err error, msgs ...string) {
	s := strings.Join(msgs, " ") + ":" + err.Error()
	log.Print(s)
	os.Exit(1)
}

func (g *Generator) fail(msgs ...string) {
	s := strings.Join(msgs, " ")
	log.Print(s)
	os.Exit(1)
}

func (g *Generator) ProtoFileBaseName(name string) string {
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}
	return name
}

func (g *Generator) generate(maker Maker, request *plugin.CodeGeneratorRequest) (*plugin.CodeGeneratorResponse, error) {
	response := new(plugin.CodeGeneratorResponse)
	for _, protoFile := range request.ProtoFile {
		if len(protoFile.GetService()) <= 0 {
			continue
		}
		file, err := maker.Make(protoFile)
		if err != nil {
			return response, err
		}
		response.File = append(response.File, file)
	}
	return response, nil
}

func (g *Generator) Generate(maker Maker) {
	input, err := ioutil.ReadAll(g.reader)
	if err != nil {
		g.error(err, "reading input")
	}

	request := g.Request
	if err := proto.Unmarshal(input, request); err != nil {
		g.error(err, "parsing input proto")
	}

	if len(request.FileToGenerate) == 0 {
		g.fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())
	g.WrapTypes()
	g.SetPackageNames()
	g.BuildTypeNameMap()
	g.GenerateAllFiles()
	g.Reset()

	response, err := g.generate(maker, request)
	if err != nil {
		g.error(err, "failed to generate files from proto")
	}

	output, err := proto.Marshal(response)
	if err != nil {
		g.error(err, "failed to marshal output proto")
	}
	_, err = g.writer.Write(output)
	if err != nil {
		g.error(err, "failed to write output proto")
	}
}
