# go.hexagonal-architecture

- each service communicate with `port` interface
- each service completely de-coupling

![hexagonal.png](hexagonal.png)

## start deps

```shell
docker-compose up 
```

## stop deps

```shell
docker-compose down
```

## generate proto

```shell
protoc --go_out=internal/adapters/framework/left/grpc --go-grpc_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/*.proto
```

## run

```shell
go run cmd/main.go
```