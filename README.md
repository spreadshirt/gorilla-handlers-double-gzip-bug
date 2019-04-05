# Minimal example for double gzip by gorilla/handlers.CompressHandler

The `CompressHandler` of [gorilla/handlers][1] gzips responses twice if they're already gzipped, i.e. it ignores the `Content-Encoding` header.

## How to reproduce

Requirements:

- go 1.12+

```sh
$ go run main.go &
$ curl -H 'Accept-Encoding: gzip' localhost:9999 | gunzip
# prints garbage
$ curl -H 'Accept-Encoding: gzip' localhost:9999 | gunzip | gunzip
# works fine
```

A working implementation is served on `localhost:19999`.

[1]: https://github.com/gorilla/handlers/commit/e1b2144f2167de0e1042d1d35e5cba5119d4fb5d
