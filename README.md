# go-gelf-logger

## Usage

Go get it: 
`go get -u github.com/imm-llc/go-gelf-logger`

**Note:** You may need to run a `go env -w GOPRIVATE=github.com/imm-llc`

Create your configuration using an instantiation of `gologger.LoggerConfig`

Initialize your logger by passing `gologger.InitLogger()` a reference to the `gologger.LoggerConfig` you created

Logging is done via an interface called `GrayLogger`, which provides three functions: `Info()`, `Warning()`, and `Error()`

Each function requires you pass it a `gologger.LogItems` struct

You can append extra log fields ad hoc as such

```go
extraFields := map[string]interface{}{"foo": "bar"}
extraFields["example"] = "whizbang"
```

## Private Module Repos in Pipelines.

You can use `go mod vendor` to ease private module access in pipelines

Save us, [Medium](https://medium.com/cloud-native-the-gathering/go-modules-with-private-git-repositories-dfe795068db4)