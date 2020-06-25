# go-gelf-logger

## Usage

Go get it: 
`go get -u github.com/imm-llc/go-gelf-logger`

**Note:** You may need to run a `go env -w GOPRIVATE=github.com/imm-llc`

Create your configuration using an instantiation of `gologger.LoggerConfig`

Initialize your logger by passing `gologger.InitLogger()` a reference to the `gologger.LoggerConfig` you created

After that, you're able to call the `Log`, `Warning`, and `Error` functions.

All of the functions require a short message, full message, and extra fields arguments. The first two are strings, extra fields is `map[string]interface{}`. See `example/main.go` for reference. The `Error` function requires you pass it an `error` that providers an `Error()` method. That error is used to build the `ErrorMessage` extra field field