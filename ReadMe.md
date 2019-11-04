# Go Service Template

Basic setup for service prototyping in Alchera Inc.

* gRPC Echo with TLS (Self-signed)

### References

* https://github.com/denji/golang-tls
* https://stackoverflow.com/questions/22666163/golang-tls-with-selfsigned-certificate
* https://bbengfort.github.io/programmer/2017/03/03/secure-grpc.html
* https://www.prakharsrivastav.com/posts/from-http-to-https-using-go/

## How To

### Setup

See [.travis.yml](./.travis.yml) for entire steps.

#### Go Path

```ps1
$env:GOPATH=Get-Location
```

```bash
export GOPATH=$(pwd)
```

#### gRPC

```
go get google.golang.org/grpc
```

#### Protocol Buffer

```bash
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/protoc-gen-go
```

#### Key/Cert Generation

The following scripts requires `openssl` tool.

```
bash ./scripts/openssl_key_gen.sh;
bash ./scripts/openssl_cert_selfsign.sh;
```

### Test

```
go test -cover -v .
```
