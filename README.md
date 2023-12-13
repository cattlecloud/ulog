# ulog [![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

ulog is a simple and efficient level logging library for Go

### Getting Started

The `ulog` package can be added to a project by running:

```shell
go get noxide.lol/go/ulog@latest
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

The `noxide.lol/go/ulog` module is open source under the [BSD](LICENSE) license.
