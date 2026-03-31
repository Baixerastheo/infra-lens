package parser

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Resource struct {
	Type       string
	Name       string
	Attributes map[string]hcl.Expression
	Blocks     map[string][]map[string]hcl.Expression
	File       string
}

func Parse(dir string) ([]Resource, error) {
	var resources []Resource

	files, err := filepath.Glob(filepath.Join(dir, "*.tf"))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		src, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		parsed, diags := hclsyntax.ParseConfig(src, file, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			return nil, diags
		}

		body := parsed.Body.(*hclsyntax.Body)
		for _, block := range body.Blocks {
			if block.Type != "resource" {
				continue
			}

			attrs := make(map[string]hcl.Expression)
			for name, attr := range block.Body.Attributes {
				attrs[name] = attr.Expr
			}

			blocks := make(map[string][]map[string]hcl.Expression)
			for _, nestedBlock := range block.Body.Blocks {
				nestedAttrs := make(map[string]hcl.Expression)
				for name, attr := range nestedBlock.Body.Attributes {
					nestedAttrs[name] = attr.Expr
				}
				blocks[nestedBlock.Type] = append(blocks[nestedBlock.Type], nestedAttrs)
			}

			resources = append(resources, Resource{
				Type:       block.Labels[0],
				Name:       block.Labels[1],
				Attributes: attrs,
				Blocks:     blocks,
				File:       file,
			})
		}
	}

	return resources, nil
}
