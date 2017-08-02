#!/bin/bash
rm example/channel_int_map_gen.go
go-bindata -pkg mapgen map.go.tmpl
go install github.com/s-kirby/mapgen/cmd/mapgen 
go generate
mapgen -p example -f example/channel_int_map_gen.go Channel/int

