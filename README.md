> **Note**
> As of Go 1.18, just use generics.

---

# mapgen

A tool which generates thread safe maps for Go.

Features:

- Supports any key/value pair supported by Go's native maps
- Allows complex operations via `Lock()` and `Unlock()`
- Generated code conforms to `golint` and `gofmt`
- Allows custom types
- Sensible default file name and map name
- Optionally generates using `sync.RWMutex`

Generated example located in `examples/`

## Install

`go get -u github.com/s-kirby/mapgen/cmd/mapgen`

## Usage

Create string -> int map:

```bash
$ mapgen string/int
Wrote string_int_map_gen.go
```

Create string -> *bytes.Buffer map using a read-write mutex:

```bash
$ mapgen --rwmu string/*bytes.Buffer
Wrote string_buffer_map_gen.go
```

Help:

```
usage: mapgen [<flags>] <keyvalue types> [<tname>]

Flags:
      --help       Show context-sensitive help (also try --help-long and --help-man).
  -p, --pkg="."    package name
  -v, --verbose    highly descriptive output
  -f, --file=FILE  output file name
      --rwmu       Use RWMutex

Args:
  <keyvalue types>  Key and value types, e.g `string/int`
  [<tname>]         name of generated type
```
