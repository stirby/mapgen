# mapgen

A tool which generates thread safe maps for Go.

Features: 
    - Supports any key/value pair supported by Go's native maps
    - Allows complex operations via `Lock()` and `Unlock()`
    - Generated code conforms to `golint` and `gofmt`
    - Allows custom types

## Example
Generated example located in `examples/`

## Usage

```bash
$ mapgen string/int
Wrote string_int_map_gen.go
```