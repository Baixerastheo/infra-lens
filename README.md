# infra-lens 🔍

> Audit your Terraform infrastructure for security issues, cost inefficiencies, and technical debt — straight from your terminal.

---

## What it does

`infra-lens` scans your Terraform files and surfaces problems before they hit production.

```bash
$ infra-lens scan ./infra

[ERROR] aws_security_group.web      → port 22 open on 0.0.0.0/0 — critical SSH exposure
[WARN]  aws_instance.api            → instance type t2.micro is deprecated, consider t3.micro (-15% cost)
[WARN]  aws_db_instance.main        → backup_retention_period = 0, no backup configured
[INFO]  aws_s3_bucket.assets        → versioning disabled, risk of data loss

─────────────────────────────────────────
  1 error · 2 warnings · 1 info
  Report saved → ./infra-lens-report.html
```

---

## Installation

```bash
# Clone the repo
git clone https://github.com/Baixerastheo/infra-lens
cd infra-lens

# Build
go build -o infra-lens .

# (optional) Move to PATH
mv infra-lens /usr/local/bin/
```

**Requirements:** Go 1.21+

---

## Usage

```bash
# Scan a directory
infra-lens scan ./infra

# Filter by severity
infra-lens scan ./infra --severity error

# Export as JSON (useful for CI/CD)
infra-lens scan ./infra --output json

# Show version
infra-lens version
```

---

## Rules

`infra-lens` checks three categories of issues:

### 🔴 Security
| Rule | Description |
|------|-------------|
| `open-ssh` | Port 22 exposed to `0.0.0.0/0` |
| `open-rdp` | Port 3389 exposed to `0.0.0.0/0` |
| `no-encryption` | Storage resources without encryption enabled |
| `wildcard-iam` | IAM policies with `*` actions |

### 🟡 Cost
| Rule | Description |
|------|-------------|
| `deprecated-instance` | EC2 instance types from previous generation |
| `missing-tags` | Resources without cost allocation tags |
| `oversized-db` | RDS instances potentially oversized for workload |

### 🔵 Technical Debt
| Rule | Description |
|------|-------------|
| `no-backup` | RDS instances with backup retention disabled |
| `no-versioning` | S3 buckets without versioning enabled |
| `hardcoded-values` | Hardcoded secrets or IPs instead of variables |

---

## CI/CD Integration

Add `infra-lens` to your GitHub Actions pipeline to audit infrastructure on every pull request:

```yaml
# .github/workflows/infra-audit.yml
name: Infra Audit

on: [pull_request]

jobs:
  audit:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - run: go build -o infra-lens .
      - run: ./infra-lens scan ./infra --severity error
```

---

## Project Structure

```
infra-lens/
├── cmd/
│   ├── root.go          # CLI entry point (Cobra)
│   └── scan.go          # scan command
├── internal/
│   ├── parser/
│   │   └── hcl.go       # Terraform file parser
│   ├── rules/
│   │   ├── engine.go    # Rule engine
│   │   ├── security.go  # Security rules
│   │   ├── cost.go      # Cost rules
│   │   └── debt.go      # Technical debt rules
│   ├── reporter/
│   │   ├── console.go   # Terminal output
│   │   └── html.go      # HTML report
│   └── models/
│       └── finding.go   # Shared structs
├── testdata/
│   └── main.tf          # Sample Terraform files for testing
├── main.go
└── README.md
```

---

## Roadmap

- [x] HCL parser
- [x] Rule engine
- [x] Console reporter
- [ ] HTML report export
- [ ] JSON output for CI/CD
- [ ] Infracost API integration
- [ ] Custom rules via config file
- [ ] Kubernetes YAML support

---

## Built with

- [Go](https://golang.org/) — because one binary is better than a node_modules folder
- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [HCL v2](https://github.com/hashicorp/hcl) — Terraform file parser
- [fatih/color](https://github.com/fatih/color) — terminal colors

---

## License

MIT — do whatever you want with it.
