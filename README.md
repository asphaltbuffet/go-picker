# Go-Picker

[![Common Changelog](https://common-changelog.org/badge.svg)](https://common-changelog.org)
[![GitHub Release](https://img.shields.io/github/v/release/asphaltbuffet/go-picker)](https://github.com/asphaltbuffet/go-picker/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/asphaltbuffet/go-picker.svg)](https://pkg.go.dev/github.com/asphaltbuffet/go-picker)
[![go.mod](https://img.shields.io/github/go-mod/go-version/asphaltbuffet/go-picker)](go.mod)
[![LICENSE](https://img.shields.io/github/license/asphaltbuffet/go-picker)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/asphaltbuffet/go-picker/build.yml?branch=main)](https://github.com/asphaltbuffet/go-picker/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/asphaltbuffet/go-picker)](https://goreportcard.com/report/github.com/asphaltbuffet/go-picker)
[![Codecov](https://codecov.io/gh/asphaltbuffet/go-picker/branch/main/graph/badge.svg)](https://codecov.io/gh/asphaltbuffet/go-picker)

⭐ `Star` this repository if you find it valuable and worth maintaining.

👁 `Watch` this repository to get notified about new releases, issues, etc.

## Description

go-picker picks one-to-many items from a list and returns them to you. It is a convenient wrapper for randomly picking elements. There are options to force the selected items to be unique and/or perform several selections based on your criteria.

## Usage

```sh
go-picker yellow red blue green --pick=2 --repeat=2 --unique=true
```

## Build

### Terminal

- `make all` - execute the build pipeline with all steps.
- `make dev` - faster execution of the build pipeline that skips some steps
- `make help` - print help for the [Make targets](Makefile).

### Visual Studio Code

`F1` → `Tasks: Run Build Task (Ctrl+Shift+B or ⇧⌘B)` to choose a build pipeline.

## Contributing

Feel free to create an issue or propose a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
