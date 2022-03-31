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

- `make` in repository directory
- Executables will be available in `build/` dir