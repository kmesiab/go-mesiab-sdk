# 🛠️ go-mesiab-sdk

An SDK for my (@kmesiab's) most commonly used stuff, like loggers,
configuration objects routers, middleware, validators, templates, etc.

## Table of Contents

1. [Badges](#badges) 📛
2. [Makefile](#makefile-) 🚀
3. [Linting](#linting-)  🧹
4. [Github Actions](#github-actions-) 👷🏼‍ <!-- markdownlint-disable-line MD051 --><!-- markdownlint-disable-line MD013 -->
5. [Config](#config-) ⚙️ <!-- markdownlint-disable-line MD051 -->

## Badges

<div align="center"><!-- markdownlint-disable-line MD033 -->

![Build and Test](https://github.com/kmesiab/go-mesiab-sdk/actions/workflows/go-build.yml/badge.svg)
![Golang Lint](https://github.com/kmesiab/go-mesiab-sdk/actions/workflows/go-lint.yml/badge.svg)
![Link Markdown](https://github.com/kmesiab/go-mesiab-sdk/actions/workflows/markdown-lint.yml/badge.svg)

![Logo](./assets/logo-transparent-small.png)

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)

</div>

## Makefile 🚀

The include makefile template already includes a lot of
targets. The most important ones are:

- `make build` - builds the project
- `make test` - runs the tests
- `make lint` - runs the linters. This includes `golangci-lint` and
`markdownlint.` You can also use `make lint-go`

## Linting 🧹

The makefile and GitHub actions are configured to run two linters,
the `golangci-lint` and `markdownlint`

🌐 [Markdown Lint](https://github.com/DavidAnson/markdownlint) -
The markdown lint project's homepage

🌐 [Golangci Lint](https://github.com/golangci/golangci-lint) -
The golang linter's project homepage

## Github Actions 👷🏼‍♂️

`readme.yml`
`go.yml`

Two Github actions are configured. One for linting the project and one
for linting the markdown.

## Config ⚙️

A simple configuration package is included that follows a chainable format:

```go
package main

import sdk "github.com/kmesiab/go-mesiab-sdk"

// Supportes format strings
sdk.Logf("An error occurred: %s", err).Debug()

// Supports packaging errors
sdk.Log("An error occurred").AddError(err).Error()

// Supports adding key value pairs
sdk.Log("An error occurred").Add("key", "value").Fatal()
```
