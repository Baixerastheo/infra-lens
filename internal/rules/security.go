package rules

import (
	"github.com/Baixerastheo/infra-lens/internal/models"
	"github.com/Baixerastheo/infra-lens/internal/parser"
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type OpenSSHRule struct{}

func (r OpenSSHRule) Check(resource parser.Resource) []models.Finding {
	var findings []models.Finding

	if resource.Type != "aws_security_group" {
		return findings
	}

	ingressBlocks, hasIngress := resource.Blocks["ingress"]
	if !hasIngress {
		return findings
	}

	for _, ingress := range ingressBlocks {
		fromPort, hasFromPort := ingress["from_port"]
		cidrBlocks, hasCidr := ingress["cidr_blocks"]

		if !hasFromPort || !hasCidr {
			continue
		}

		portVal, diags := fromPort.Value(&hcl.EvalContext{})
		if diags.HasErrors() || portVal.Type() != cty.Number {
			continue
		}

		port, _ := portVal.AsBigFloat().Int64()

		cidrVal, diags := cidrBlocks.Value(&hcl.EvalContext{})

		hasOpenCidr := false
		for it := cidrVal.ElementIterator(); it.Next(); {
			_, v := it.Element()
			if v.Type() == cty.String && v.AsString() == "0.0.0.0/0" {
				hasOpenCidr = true
				break
			}
		}

		if port == 22 && hasOpenCidr {
			findings = append(findings, models.Finding{
				Severity: models.SeverityError,
				Resource: resource.Type + "." + resource.Name,
				Message:  "port 22 open on 0.0.0.0/0 — critical SSH exposure",
				Rule:     "open-ssh",
				File:     resource.File,
			})
		}
	}

	return findings
}
