package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	tfjson "github.com/hashicorp/terraform-json"
	tfaddr "github.com/hashicorp/terraform-registry-address"
	"github.com/hashicorp/terraform-schema/earlydecoder"
	"github.com/hashicorp/terraform-schema/module"
	tfschema "github.com/hashicorp/terraform-schema/schema"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

const vpcid = "myvpcid"
const tagid = "mydemotag"
func main() {

	s := mergeSchema()
	p := hclparse.NewParser()
	f, diag := p.ParseHCLFile("test/ec2.tf")
	if diag.HasErrors() {
		fmt.Println("Couldn't parse file")
	}
	hclSchema := s.ToHCLSchema()
	c, diag := f.Body.Content(hclSchema)

	for _, b := range c.Blocks {
		fmt.Printf("%v\n", b.Type)
		if contains(b.Labels, "aws_instance") {
			parseEC2(b, hclSchema)
		}
	}
	// create new file on system
	tfFile, err := os.Create("outputfile.tf")
	if err != nil {
		fmt.Println(err)
		return
	}
	// initialize the body of the new file object

	fmt.Printf("%s", f.Bytes)
	_, err = tfFile.Write(f.Bytes)
	if err != nil {
		fmt.Printf("Error writing file %v", err)
	}
}

func parseEC2(b *hcl.Block, s *hcl.BodySchema) {
	fmt.Printf("body: %v\n", b.Body)
	content, otherbody, diag := b.Body.PartialContent(s)
	fmt.Printf("%v", content)
	fmt.Printf("%v", otherbody)
	if diag.HasErrors() {
		fmt.Printf("aws errors: %v", diag.Error())
	}
	//attr := otherbody.Attributes

	//fmt.Printf("attributes: %v\n", attr)
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func mergeSchema() *schema.BodySchema {
	coreSchema := tfschema.UniversalCoreModuleSchema()
	v := version.Must(version.NewVersion("0.12.0"))
	sm := tfschema.NewSchemaMerger(coreSchema)

	ps := &tfjson.ProviderSchemas{}
	b, err := ioutil.ReadFile("test/provider.json")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	err = json.Unmarshal(b, ps)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	sr := &testJsonSchemaReader{
		ps: ps,
		migrations: map[tfaddr.Provider]tfaddr.Provider{
			// the builtin provider doesn't have entry in required_providers
			tfaddr.NewLegacyProvider("terraform"): tfaddr.NewBuiltInProvider("terraform"),
		},
	}
	sm.SetSchemaReader(sr)
	sm.SetTerraformVersion(v)
	mergedSchema, err := sm.SchemaForModule(testModuleMeta("test/ec2.tf"))

	return mergedSchema
}

type testJsonSchemaReader struct {
	ps          *tfjson.ProviderSchemas
	useTypeOnly bool
	migrations  map[tfaddr.Provider]tfaddr.Provider
}

func (r *testJsonSchemaReader) ProviderSchema(_ string, pAddr tfaddr.Provider, _ version.Constraints) (*tfschema.ProviderSchema, error) {
	if newAddr, ok := r.migrations[pAddr]; ok {
		pAddr = newAddr
	}

	addr := pAddr.String()
	if r.useTypeOnly {
		addr = pAddr.Type
	}

	jsonSchema, ok := r.ps.Schemas[addr]
	if !ok {
		return nil, fmt.Errorf("%s: schema not found", pAddr.String())
	}

	return tfschema.ProviderSchemaFromJson(jsonSchema, pAddr), nil
}
func testModuleMeta(path string) *module.Meta {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	filename := filepath.Base(path)

	f, diags := hclsyntax.ParseConfig(b, filename, hcl.InitialPos)
	if len(diags) > 0 {
		fmt.Printf("err: %v", diags)
	}
	meta, diags := earlydecoder.LoadModule("testdata", map[string]*hcl.File{
		filename: f,
	})
	if diags.HasErrors() {
		fmt.Printf("err: %v", diags)
	}
	return meta
}
