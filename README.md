# DNS Extension for Fn

A DNS extension to support nice domains on [Fn](https://fnproject.io).

## Usage

Add this as middleware in your main.go:

```go
funcServer.AddRootMiddleware(&dns.Middleware{})
```

Point a wildcard domain at your hosted server.

## TODO

* [ ] support per app custom domains

## Contributing

To test this locally, add a line to `/etc/hosts`

```txt
127.0.0.1 myapp.local.com
```

Start server with the DNS middleware:

```sh
cd main
export API_HOST=localhost,myapp.local.com
go build && ./main
```

In another window:

```sh
fn init --runtime go gofunc
cd gofunc
fn deploy --app myapp --local
curl http://myapp.local.com:8080/gofunc
```
