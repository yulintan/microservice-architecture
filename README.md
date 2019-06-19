# Description
This monorepo demo project shows you how gRPC is used for restful API as well as internal communication between each microservice.

##### Why gRPC
- Faster
- Language agnostic
- Exchange messages in binary
- can be Auto generated
- Reusable for internal communication

##### Why Restful
- Easy to understand
- Http is already widely used

##### Why microservice
- Easy deployment
- Decoupled
- Scalable
- Distributed

# Prerequisite
- Run `go run main.go` under product-api folder to host product gRPC & restful server
- Run `go run main.go` under shop-api folder to host shop gRPC & restful server

(To be continue)



![Microservice](https://github.com/yulintan/microservice-architecture/blob/master/architecture.png)
