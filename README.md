# go-gelf-logger

## Usage

Go get it: 
`go get -u github.com/imm-llc/go-gelf-logger`

Create your configuration using an instantiation of `gologger.LoggerConfig`

Initialize your logger by passing `gologger.InitLogger()` a reference to the `gologger.LoggerConfig` you created

Logging is done via an interface called `GrayLogger`, which provides three functions: `Info()`, `Warning()`, and `Error()`

Each function requires you pass it a `gologger.LogItems` struct

You can append extra log fields ad hoc as such

```go
extraFields := map[string]interface{}{"foo": "bar"}
extraFields["example"] = "whizbang"
```

You can find a full example in the `example/` directory.