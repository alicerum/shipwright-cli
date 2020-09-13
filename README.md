`shp` (PoC)
-----------

This is a proof-of-concept implementation of a command-line client for
[Shipwright's Build](shipwrightbuild) operator. It splits the concerns in `cmd` and `pkg`, where
`cmd` holds all the interactions with command-line, while `pkg` works as a specialized API client.

To install it run:

```sh
go get github.com/otaviof/shp/cmd/shp
```

Or clone the repository, and:

```sh
make
```

For instance:

```sh
output/shp build-run \
    --buildref-name="test" \
    --buildref-apiversion="v1" \
    --sa-name="sa" \
    --sa-generate \
    --timeout="10s" \
    --output-image="image" \
    --output-credentials="secret" \
    create test
```

[shipwrightbuild]: https://github.com/shipwright-io/build/
