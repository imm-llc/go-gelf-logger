# go-gelf-logger

## Usage

Go get it: 
`go get -u github.com/imm-llc/go-gelf-logger`

### v0.2.2 and Earlier

Create your configuration using an instantiation of `gologger.LoggerConfig`

Initialize your logger by passing `gologger.InitLogger()` a reference to the `gologger.LoggerConfig` you created

Logging is done via an interface called `GrayLogger`, which provides three functions: `Info()`, `Warning()`, and `Error()`

Each function requires you pass it a `gologger.LogItems` struct

You can append extra log fields ad hoc as such

```go
extraFields := map[string]interface{}{"foo": "bar"}
extraFields["example"] = "whizbang"
```

### v0.3.0 and Later

Create your configuration using an instantiation of `gologger.LoggerConfig`

You may either initialize a UDP or TCP logger.

UDP: pass `gologger.InitUDPLogger()` a reference to the `gologger.LoggerConfig` you created

TCP: pass `gologger.InitTCPLogger()` a reference to the `gologger.LoggerConfig` you created

Logging is done via an interface called `GrayLogger`, which provides three functions: `Info()`, `Warning()`, and `Error()`

Each function requires you pass it a `gologger.LogItems` struct

You can append extra log fields ad hoc as such

```go
extraFields := map[string]interface{}{"foo": "bar"}
extraFields["example"] = "whizbang"
```

You can find a full example in the `example/` directory.

## A Note on Versions

Version 0.3.0 is a major overhaul of this module. Code using v0.2.2 and earlier is not compatable with v0.3.0 and later.

## Testing

To set up Graylog locally, you can use `docker-compose` in `./tests/config`. 

Once all of the containers are running, run `./tests/config/bootstrap` to create a TCP and UDP input, each running on port 11111.

Once the inputs are running, run `go test -v ./gologger`.