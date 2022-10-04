<p align="center">
  <img
    height="300"
    src="./.github/golang-and-grpc-banner.png"
    alt="Golang + gRPC"
  />
</p>

# Golang + gRPC

## Overview

This application implements a gRPC server with a service to create and list sample products

- [x] gRPC Server
- [x] gRPC Service
- [x] gRPC .proto
- [x] Evans Client gRPC
- [x] HTTP Server
- [x] Docker + Docker Compose

---

## Setup

### Install dependencies

```bash
$ go mod tidy
```

### Run app

The environment were configured using Docker Compose, to start the environment you must run:

```bash
$ docker-compose up
```

Access the application container:

```bash
$ docker exec -it go-grpc-app /bin/bash
```

Now start app:

```bash
$ go run main.go
```

### Use the Evans gRPC client

Inside the application container execute:

```bash
$ evans -r repl
```

There is only 2 RPC in the `ProductService` service.

To use RPC `CreateProduct` run:

```txt
goGrpc.ProductService@127.0.0.1:50051> call CreateProduct
```

To use RPC `ListProducts` run:

```txt
goGrpc.ProductService@127.0.0.1:50051> call ListProducts
```

![evans-grpc-client](.github/evans-grpc-client.png)

### Protobuf

After creating the .proto file, it is necessary generates Go language bindings of services in protobuf definition files for gRPC.

The necessary files have already been generated and are available in the folder `grpc/pb`.

But if you need to generate it again after some modification in the `.proto` file, a [script](Makefile) was created to generate it automatically, just enter the application container and run the command at the root of the project:

```bash
$ docker exec -it go-grpc-app /bin/bash
$ make gen
```
