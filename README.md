# DNS Extension for Fn

A DNS extension to support nice domains on [Fn](https://fnproject.io).

## Usage

First build a custom Fn image by adding this to your ext.yaml:

`github.com/fnproject/ext-dns`

See: https://github.com/fnproject/fn/blob/master/docs/operating/extending.md for how to build a custom image.

Then add a wildcard domain to your DNS provider that points to your Fn Server/Cluster:

```txt
*.mydomain.com -> myservers.fn.domain.com
```

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
go run main.go
```

In another window:

```sh
fn init --runtime go gofunc
cd gofunc
fn deploy --app myapp --local
curl http://myapp.local.com:8080/gofunc
```
