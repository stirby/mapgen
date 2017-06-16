package mapgen

import "io"
import "text/template"

//go:generate go-bindata -pkg mapgen map.go.tmpl

var mapTmpl = template.Must(template.New("map").Parse(string(MustAsset("map.go.tmpl"))))

// Generate creates a map from provided options
func Generate(params Params, w io.Writer) error {
	return mapTmpl.Execute(w, params)
}
