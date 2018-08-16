package main

import "github.com/qneyrat/go-grpc-endpoints/protoc-gen-goendpoints/gen"

func main() {
	g := gen.New()
	g.Generate()
}
