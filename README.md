<h1 align="center">gRPC with Go</h1>
<p align="center">This is a POC (proof of concept) to understand better the behavior of gRPC with Go Lang.</p>


#### Technologies:

* Go;
* gRPC;
* Protocol Buffers (https://developers.google.com/protocol-buffers)

#### How to get started:

```console
#.profile
export GOROOT=/usr/local/go
export GOPATH=$HOME/projects/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
```

```console
bscpaz@2am:/$ mkdir pb
bscpaz@2am:/$ mkdir proto 
bscpaz@2am:/$ sudo chown bscpaz:bscpaz pb proto

#it creates a new go module.
bscpaz@2am:/$ go mod init github.com/bscpaz/poc-grpc-go

bscpaz@2am:/$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
bscpaz@2am:/$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

#it generates stubs services
bscpaz@2am:/$ protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```


 

