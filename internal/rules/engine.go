package rules

import (
	"github.com/Baixerastheo/infra-lens/internal/models"
	"github.com/Baixerastheo/infra-lens/internal/parser"
)

type Rule interface {
	Check(resource parser.Resource) []models.Finding
}

type Engine struct {
	rules []Rule
}

func NewEngine() *Engine {
	return &Engine{
		rules: []Rule{
			OpenSSHRule{},
		},
	}
}

func (e *Engine) Run(resources []parser.Resource) []models.Finding {
	var findings []models.Finding

	for _, resource := range resources {
		for _, rule := range e.rules {
			findings = append(findings, rule.Check(resource)...)
		}
	}

	return findings
}
