# ulog

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/ulog.svg)](https://pkg.go.dev/cattlecloud.net/go/ulog)
[![License](https://img.shields.io/github/license/cattlecloud/ulog?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/ulog/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/ulog/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/ulog/actions/workflows/ci.yaml)

`ulog` is a simple and efficient level logging library for Go

### Getting Started

The `ulog` package can be added to a project with `go get`.

```shell
go get cattlecloud.net/go/ulog@latest
```

### Examples

```go
log := ulog.New("my/component")
log.E.Fmt("the level is %s", "error")
log.W.Fmt("the level is %s", "warn")
log.I.Fmt("the level is %s", "info")
log.D.Fmt("the level is %s", "debug")
log.T.Fmt("the level is %s", "trace")
```

### License

The `cattlecloud.net/go/ulog` module is open source under the [BSD-3-Clause](LICENSE) license.
