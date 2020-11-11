# Welcome to Go GRPC Protobuf Example!

Hi! This repo is prepared to show you how to use **protocol buffers (protobuf)** and **grpc** in **golang** together.

## Download and install the requirements

### Protocol Buffers
Download the latest prebuilt version here:
https://github.com/protocolbuffers/protobuf/releases/latest

Extract the compressed file. Then;
 - Copy contents of **bin** folder to **/usr/local/bin**
 - Copy **google** folder in **include** to **/usr/local/include** 
  
### Go Protobuf and GRPC Package

Execute these commands:  
```
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
  
## Create Proto File
Create an empty project folder (let's say "myproject") and start writing your proto file. Here is an example:

    syntax = "proto3";
    
    option go_package = ".;pb";
    
    message  Person {
    string  name = 1 [json_name="Name"];
    int32  age = 2 [json_name="Age"];
    string  address = 3 [json_name="Address"];
    }
    
    service  PersonService {
    rpc  Add (PersonRequest) returns (PersonResponse);
    }
    
    // The request message containing the user's name.
    message  PersonRequest {
    string  name = 1;
    }
    
    // The response message containing the greetings
    message  PersonResponse {
    string  name = 1;
    }

In **myproject** folder create a folder named **pb**. This folder will contain our generated protobuf and grpc files. After that run this command in **myproject** folder.

    protoc --go_out=pb --go_opt=paths=source_relative \
    	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    	./command.proto

If everything works fine, then you will have two files in **pb** folder.  
**person.pb.go** // protobuf  
**person_grpc.pb.go** // grpc client and server functions

## Work with examples
After the steps above, you can work on client and server examples in this repo. I hope it helps you.

### Resources
https://grpc.io/docs/languages/go/quickstart/  
https://developers.google.com/protocol-buffers/docs/overview