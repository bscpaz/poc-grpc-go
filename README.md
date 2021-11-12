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

#it generates stubs services
bscpaz@2am:/$ protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```


 

