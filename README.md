<h1 align="center">gRPC with Go</h1>
<p align="center">This is a POC (proof of concept) to understand better the behavior of gRPC with Go Lang.</p>


#### Technologies:

* Go (https://golang.org/);
  * _Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of **multicore** and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language._
  * https://go.dev/ref/spec


* gRPC (https://grpc.io/);
  * _In gRPC, a client application can directly call a method on a server application on a different machine as if it were a local object, making it easier for you to create distributed applications and services._
  <img src="https://user-images.githubusercontent.com/9732874/142060820-6dac8f12-1b1e-4c53-9e30-d3f856004557.png" width="700" align="center"/>

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

#Intall protocol buffers on machine.
bscpaz@2am:/$ go get -u google.golang.org/grpc
bscpaz@2am:/$ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
bscpaz@2am:/$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

#it generates stubs in Go language
bscpaz@2am:/$ protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```

#### How to start gRPC server (stream server):

```console
bscpaz@2am:/$ go run cmd/server/server.go
```

#### How to call gRPC server:

```console
bscpaz@2am:/$ go run cmd/client/client.go
```

#### How to test gRPC calls:
Install and use the Evans project:

* https://github.com/ktr0731/evans

```console
#After extract evans program into "/usr/local/go/bin" folder, add evans into .profile file:
export EVANS=$GOBIN/evans
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN:$EVANS

#Create needed folder at user's home folder.
bscpaz@2am:/$ sudo mkdir .cache/evans
bscpaz@2am:/$ sudo chown -R bscpaz:bscpaz .cache/

#From project's folder.
bscpaz@2am:/$ evans -r repl --host localhost --port 50051
pb.UserService@localhost:50051> service UserService
pb.UserService@localhost:50051> call AddUser
id (TYPE_INT32) => 0
name (TYPE_STRING) => Bruno Paz
email (TYPE_STRING) => soujava@gmail.com

#Below is the result of client call to server. Note the value of "id":
{
  "email": "soujava@gmail.com",
  "id": 123456,
  "name": "Bruno Paz"
}

pb.UserService@localhost:50051> exit
```


<hr>
<h4 align="center">Known issues</h4>

```console
Issue:
  could not import "google.golang.org/grpc"
Solution:
  bscpaz@2am:/$ go get -u google.golang.org/grpc
``` 
```console
Issue:
    bscpaz@2am:$ evans -r repl --host localhost --port 50051
    evans: failed to run REPL mode: failed to instantiate a new spec: failed to instantiate the spec: 
    failed to list packages by gRPC reflection: failed to list services from reflection enabled gRPC server: 
    rpc error: code = Unimplemented desc = unknown service grpc.reflection.v1alpha.ServerReflection
Solution:
    Add reflection mode into code: 
      reflection.Register(grpcServer)
``` 
```console
Issue:
  Command 'go' not found, but can be installed with:
  or
  evans: command not found
Solution:
  You've opened a new terminal and it requires a new refresh on .profile file.
  bscpaz@2am:/$ cd ~
  bscpaz@2am:/$ source .profile
``` 
```console
Issue:
  protoc-gen-go: program not found or is not executable
 Solution:
 
``` 
