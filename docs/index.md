# actions_test

Go project demonstrating CI/CD best practices with GitHub Actions.

## Overview

Module: `github.com/mrashed/actions_test`

## Getting Started

```bash
go build ./...
go test ./...
```

## CI/CD

This project uses GitHub Actions for:

- SAST (CodeQL)
- Dependency review and vulnerability scanning (Trivy)
- Secret scanning (Gitleaks)
- Lint and format checks (golangci-lint)
- Automated release notes (git-cliff)
- SBOM generation (Syft)
- MkDocs documentation build and deploy
