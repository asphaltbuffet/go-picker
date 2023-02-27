# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.0] - 2023-02-26

### Added

- Initial project creation from template
- added `godox`, `gomnd`, `prealloc`, `maintidx`, and `wsl` linters
- ignore `dupl` and `funlen` in test files
- explicit configuration of release checksum file
- additional [VSCode](.vscode/tasks.json) tasks for building
- include [gofumpt](https://mvdan.cc/gofumpt) to project tools

### Changed

- updated [README](README.md) to be project-specific
- split build task into two variants (`all` + `dev`)
- replace use of `--rm-dist` with `--clean` in goreleaser
- replace deprecated `name-template` section in [goreleaser config](.goreleaser.yml)
- target go version 1.20
- put coverage artifacts in `bin` directory

### Removed

- replace use of tools directory and files with explicit definition in [Makefile]
- favor use of misspell linter instead of separate build step

[Unreleased]: https://github.com/asphaltbuffet/go-picker/compare/v0.0.0...HEAD
[0.0.0]: https://github.com/asphaltbuffet/go-picker/releases/tag/v0.0.0
