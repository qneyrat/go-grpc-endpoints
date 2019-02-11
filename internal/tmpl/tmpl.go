package tmpl

import (
	"bytes"
	"strings"
	"text/template"

    "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type Template struct {
	*template.Template
}

func getFuncs() template.FuncMap {
	return template.FuncMap{
		"getProtoName": func(proto descriptor.MethodDescriptorProto) string {
			return proto.GetName()
		},
		"getProtoNameToLower": func(proto descriptor.MethodDescriptorProto) string {
			return strings.ToLower(proto.GetName())
		},
		"getProtoInputType": func(proto descriptor.MethodDescriptorProto) string {
			return strings.Split(proto.GetInputType(), ".")[2]
		},
		"getProtoOutputType": func(proto descriptor.MethodDescriptorProto) string {
			return strings.Split(proto.GetOutputType(), ".")[2]
		},
		"camelCase": func(str string) string {
			return generator.CamelCase(str)
		},
	}
}

func NewTempale(buf []byte) (*Template, error) {
	t := template.New("template")
	t.Funcs(getFuncs())
	_, err := t.Parse(string(buf))
	if err != nil {
		return &Template{}, err
	}
	return &Template{ t }, nil
}

func (t Template) Eval(data interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	if err != nil {
		return nil , err
	}
	return buf, nil
}
