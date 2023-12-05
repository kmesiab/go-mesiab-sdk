# 🛠️ go-mesiab-sdk

An SDK for my (@kmesiab's) most commonly used stuff, like loggers, configuration objects
routers, middleware, validators, templates, etc.

# Table of Contents

1. [Badges](#Badges) 📛
2. [Makefile](#makefile-) 🚀
3. [Linting](#linting-)  🧹
4. [Github Actions](#github-actions-%EF%B8%8F) 👷🏼‍
5. [Config](#config-%EF%B8%8F) ⚙️

## Badges

![Build](https://github.com/kmesiab/go-audibly/actions/workflows/go.yml/badge.svg)

![Logo](./assets/logo-small.png)

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white)
![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)

## Makefile 🚀

The include makefile template already includes a a lot of
targets. The most important ones are:

- `make build` - builds the project
- `make test` - runs the tests
- `make lint` - runs the linters. This includes `golangci-lint` and 
`markdownlint.` You can also use `make lint-go`

## Linting 🧹

The makefile and github actions are configured to run two linters,
the `golangci-lint` and `markdownlint`

🌐 [Markdown Lint](https://github.com/DavidAnson/markdownlint) - The markdown lint project's homepage

🌐 [Golangci Lint](https://github.com/golangci/golangci-lint) - The golang linter's project homepage

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
