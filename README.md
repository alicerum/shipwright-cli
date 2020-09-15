`shp` (PoC)
-----------

This is a proof-of-concept implementation of a command-line client for
[Shipwright's Build](shipwrightbuild) operator. It splits the concerns in `cmd` and `pkg`, where
`cmd` holds all the interactions with command-line, while `pkg` works as a specialized API client.

## Install

To install it run:

```sh
go get github.com/otaviof/shp/cmd/shp
```
## Build

Or clone the repository, and:

```sh
make
```

### `kubectl` Plugin

In order to compile the project as a [kubectl plugin][kubectlplugin], run:

```sh
make kubectl
install -m +x _output/kubectl-shp /usr/local/bin/
```

And then you will be able to use as `kubectl shp`, in command-line.

## Usage

The example below shows the creation of a simple `BuildRun` resource named `test`.

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
[kubectlplugin]: https://krew.sigs.k8s.io/docs/developer-guide/
