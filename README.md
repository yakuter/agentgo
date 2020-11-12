# Welcome to Agentgo!

Hi! Agentgo is a tool for making remote command executions with **golang**, **protocol buffers (protobuf)** and **grpc**. This is good way to do some operations at agents (let's say clients). This is why it is called **Agentgo**.
  
### Proto File Update
If you change anything in **command.proto** file, then you need to update auto generated codes with protobuf. To do this, you can use the command below just after you update **command.proto**.

    protoc --go_out=pb --go_opt=paths=source_relative \
    	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    	./command.proto

If everything works fine, then you will have two files in **pb** folder.  
**person.pb.go** // protobuf  
**person_grpc.pb.go** // grpc client and server functions

### Resources
https://grpc.io/docs/languages/go/quickstart/  
https://developers.google.com/protocol-buffers/docs/overview