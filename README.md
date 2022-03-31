# Simple async client-server app

## Usage

```shell
goserver [-p 8080] [-f history.msg] 
```

```shell
goclient 127.0.0.1:8080
```
Type `goserver help` or `goclient help` for help.

## User CLI

Type message and press `Enter`, it will be sent to a server and written in log file. Type `\exit` to exit session.

## Build
- **requirement:** installed `go 1.17` or higher
- Makefile for UNIX
- `make` in repository directory
- Executables will be available in `build/` dir

## Execution Policy

- Server accept new connection in endless loop
- The only way to stop it is to send SIGINT from console or SIGKILL
- If main loop is interrupted, all goroutines are serving clients will be stopped immediately
- If server disconnect, client will see `error: broken pipe` when trying to send message 