@echo off

set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o micro_lin micro.go

echo make exe files is ok!
