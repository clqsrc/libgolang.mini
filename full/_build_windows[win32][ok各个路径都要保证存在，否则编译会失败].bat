
set GOPATH=D:\gopath1.10.8

#set GOROOT=D:\go1.10.8.windows-amd64\go
set GOROOT=D:\go1.10.8.windows-386\go

rem set PATH=D:\go1.10.8.windows-amd64\go\bin;D:\new\Qt\Qt5.12.12\Tools\mingw730_64\bin
rem set PATH=D:\new\mingw64-8.1.0-sjlj\mingw64\bin;%GOROOT%\bin;%PATH%
PATH=D:\go1.10.8.windows-386\go\bin;D:\new\Qt\Qt5.12.12\Tools\mingw730_32\bin


set GOARCH=386
rem set GOARCH=amd64
set GOOS=windows
# CGO_ENABLED 目前也一定要有
set CGO_ENABLED=1

rem go build -buildmode=c-shared -o libgolang.dll http_dll.go

go build -buildmode=c-shared -o libgolang.dll



