# Docker context interactive CLI 🇺🇦
Interactive CLI tool to switch between docker contexts. Use with `-s` flag to ssh to the selected host instead of "using" it.
#### Handy use with aliases, e.g.
```bash
alias dps="docker ps -a"
alias dc="docker-compose"
alias du="docker-context-interactive-mac"
alias ds="docker-context-interactive-mac -s"
alias dd="docker context use default"
alias dl="docker context ls"
```

### To build with `docker`
```bash
docker run -it --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=amd64 golang:1.22.5 go build -v -ldflags "-s -w" -o ./bin/docker-context-interactive-linux main.go
docker run -it --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=windows -e GOARCH=amd64 golang:1.22.5 go build -v -ldflags "-s -w" -o ./bin/docker-context-interactive-win.exe main.go
docker run -it --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=darwin golang:1.22.5 go build -v -ldflags "-s -w" -o ./bin/docker-context-interactive-mac main.go
```

### To run dev with `docker`
```bash
docker run -it --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.22.5 bash
go run main.go
```
