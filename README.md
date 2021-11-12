<h1 align="center">gRPC with Go</h1>
<p align="center">This is a POC (proof of concept) to understand better the behavior of gRPC with Go Lang.</p>


#### Technologies:

* Go;
* gRPC;
* Protocol Buffers (https://developers.google.com/protocol-buffers)

#### How to get started:

```console
#it creates a new go module.
bscpaz@2am:/$ go mod init github.com/bscpaz/poc-grpc-go

#it generates stubs services
bscpaz@2am:/$ protoc --proto_path=proto proto/*.proto --go_out=pb
```


 

