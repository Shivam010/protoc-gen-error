package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
)

//go:generate protoc -I . types.proto --go_out=:.
//go:generate go build -o bin/protoc-gen-error
//go:generate protoc -I . check.proto --plugin=bin/protoc-gen-error --error_out=:.

func main() {
	pgs.Init().RegisterModule(&cMod{&pgs.ModuleBase{}}).Render()
}

type cMod struct {
	*pgs.ModuleBase
}

func (m *cMod) Name() string {
	return "Error Report"
}

func (m *cMod) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, fl := range targets {
		for _, msg := range fl.Messages() {

			content := ""

			for _, f := range msg.Fields() {

				opt := f.Descriptor().GetOptions()

				rule, _ := proto.GetExtension(opt, E_Rule)
				if rule, ok := rule.(*Rule); ok && rule != nil {
					content += fmt.Sprintf("\n* Options of the field \"%s\": %+v \n", f.Name().String(), rule.Type)
				}
			}

			m.OverwriteGeneratorFile("report.txt", content)
		}
	}
	return m.Artifacts()
}
