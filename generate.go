package mapgen

import "io"
import "text/template"
import "strings"

//go:generate go-bindata -pkg mapgen map.go.tmpl

var mapTmpl = template.Must(template.New("map").Funcs(template.FuncMap{
	"title": strings.Title,
}).Parse(string(MustAsset("map.go.tmpl"))))

// Generate creates a map from provided options
func Generate(params Params, w io.Writer) error {
	return mapTmpl.Execute(w, params)
}
