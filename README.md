# GoRootFS

GoRootFS is a minimal, modular shell written in Go, designed for Termux/Android.  
It features internal commands and a dynamic prompt  

## Features

- Internal commands: `ls`, `cd`, `pwd`, `cat`, `touch`, `mkdir`, `rm`, `cp`, `mv`, `whoami`, `echo`, `clear`
- Dynamic prompt showing current directory (`~` for home)
- Fully written in Go, statically compiled

## Build

Build statically for your architecture:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -ldflags="-s -w" -o goshell main.go
```

## Usage

```bash
./goshell
```
