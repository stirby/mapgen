# mapgen

A tool which generates thread safe maps for Go.

Features: 
    - Supports any key/value pair supported by Go's native maps
    - Allows complex operations via `Lock()` and `Unlock()`
    - Generated code conforms to `golint` and `gofmt`
    - Allows custom types

Generated example located in `examples/`

## Usage


Create string -> int map

```bash
$ mapgen string/int
Wrote string_int_map_gen.go
```

Create string -> *bytes.Buffer map

```bash
$ mapgen string/*bytes.Buffer
Wrote string_buffer_map_gen.go
```

Help:

```
usage: mapgen [<flags>] <keyvalue type>

Flags:
      --help         Show context-sensitive help (also try --help-long and --help-man).
  -p, --pkg="."      package name
  -v, --verbose      highly descriptive output
  -t, --tname=TNAME  name of generated type
  -f, --fname=FNAME  file name of generated type

Args:
  <keyvalue type>  Key value type, e.g `string/int`
```