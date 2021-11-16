<h1 align="center">gRPC with Go</h1>
<p align="center">This is a POC (proof of concept) to understand better the behavior of gRPC with Go Lang.</p>


#### Technologies:

* Go (https://golang.org/);
  * _Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of **multicore** and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language._


* gRPC (https://grpc.io/);
  * _In gRPC, a client application can directly call a method on a server application on a different machine as if it were a local object, making it easier for you to create distributed applications and services._
  * <img src="https://user-images.githubusercontent.com/9732874/142056444-3b1efa74-dc2e-4ff1-a58c-397063d63922.png" width="500"/>


* Protocol Buffers (https://developers.google.com/protocol-buffers)
  * _Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages._

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


 

