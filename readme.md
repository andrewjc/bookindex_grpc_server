### Book Index Service

This is an grpc service for indexing books. Provides a query search endpoint.

Sample usage:

```bash
grpcurl -plaintext -d '{"query": "golang"}' localhost:8080 BookIndex/Search
```

### Development

You can get started quickly by running the following:

```bash
make install
make run
```

Regenerate protobuffer files at any time with:

```bash
make proto
```

#### Requirements

- [Go](https://golang.org/doc/install)
- [Protocol Buffers](https://developers.google.com/protocol-buffers/docs/gotutorial)
- [gRPC](https://grpc.io/docs/quickstart/go/)
- [grpcurl](https://github.com/fullstorydev/grpcurl)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [golangci-lint](https://golangci-lint.run/usage/install/)
- [goreleaser](https://goreleaser.com/install/)
