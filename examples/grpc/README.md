## How to run this example

1. run grpc server

```sh
$ go run grpc/server.go
```

2. run ego server

```sh
$ go run ego/main.go
```

3. use curl command to test it

```sh
$ curl -v 'http://localhost:8052/rest/n/thinkerou'
```
