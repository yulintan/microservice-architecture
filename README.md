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
- Loosely coupled
- Scalable
- Distributed

# Prerequisite
- Run `go run main.go` under product-api folder to host product gRPC & restful server
- Run `go run main.go` under shop-api folder to host shop gRPC & restful server

# Workflow
There are two microservices (product-api and shop-api) and 2 endpoints which return dummy data
- get shop at http://localhost:8080/api/shops/1 which returns dummy shop response
```
{
    shop: {
        id: "1",
        shop_domain: "test1.shop.com",
        currency: "$",
        created_at: null,
        updated_at: null
    }
}
```

- get product at http://localhost:8081/api/products/1, which returns dummy product response
```
{
    product: {
        id: "1",
        name: "Test Name 1",
        price: "100.00",
        currency: "$", // product database doesn't store this information, we need fetch this from shop-api via rpc call
        created_at: null,
        updated_at: null
    }
}
```

In the product-api, it doesn't know anything about shop but it needs get the shop currency.

What's happening when hitting the get product endpoint:

1. Get product http request is routed to product RPC server

2. GetProductByID function is called in product-api/rpci/server.go

3. From product-api, it calls shop service via shop rpc server to get the currency

4. Return product object with correct currency


# Benifit

- OpenAPI&sdk, http router, rpc client, rpc server are all generated, we just need to focus on the business logic.
- Each function/service can be reused mutiple times. It's real Service-oriented architecture.
- Smaller code base, easier for onbording more devs.
- Easy to manage, deploy and scale up&down.

# Diagram
![Microservice](https://github.com/yulintan/microservice-architecture/blob/master/architecture.png)



