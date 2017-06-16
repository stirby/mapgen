package main

import (
	"fmt"
	"os"
	"path/filepath"
	"unicode"

	"strings"

	"os/user"

	"os/exec"

	"github.com/alecthomas/kingpin"
	"github.com/ammario/mapgen"
	"github.com/ammario/micropkg/strutil"
	"github.com/fatih/color"
	"github.com/pkg/errors"
)

func fail(msg string, args ...interface{}) {
	color.New(color.FgHiRed).Printf(msg+"\n", args...)
	os.Exit(1)
}

// failErr fails if err != nil
func failErr(err error) {
	if err != nil {
		fail("error: %+v", err)
	}
}

func gopath() string {
	path := os.Getenv("GOPATH")
	if path != "" {
		return path
	}
	user, err := user.Current()
	failErr(errors.Wrap(err, "failed to get user"))
	return user.HomeDir + "/go/"
}

// justType turns a type like `example.Channel` into just `Channel``
func justType(fullType string) (t string) {
	tokens := strings.Split(fullType, ".")
	return tokens[len(tokens)-1]
}

func main() {
	// flags
	var (
		pkgName       = kingpin.Flag("pkg", "package name").Short('p').Default(".").String()
		verbose       = kingpin.Flag("verbose", "highly descriptive output").Short('v').Bool()
		typName       = kingpin.Flag("tname", "name of generated type").Short('t').String()
		fileName      = kingpin.Flag("fname", "file name of generated type").Short('f').String()
		keyValueTypes = kingpin.Arg("keyvalue types", "Key and value types, e.g `string/int`").Required().String()

		useRwMutex = kingpin.Flag("rwmu", "Use RWMutex").Default("false").Bool()
	)
	kingpin.Parse()
	var err error

	if *pkgName == "." {
		*pkgName, err = os.Getwd()
		failErr(errors.Wrap(err, "failed to get working dir"))
		*pkgName = filepath.Base(*pkgName)
	}
	if *verbose {
		fmt.Printf("pkg: %v\n", *pkgName)
	}

	// parse basic map info
	key, value, err := mapgen.ParseKeyValueType(*keyValueTypes)
	failErr(err)
	if *verbose {
		fmt.Printf(" key: %v, value: %v\n", key, value)
	}

	simpleKey, simpleVal := justType(key), justType(value)

	// figure out reasonable type name
	if *typName == "" {
		*typName = strings.Title(simpleKey) + strings.Title(simpleVal) + "Map"
	}
	if *verbose {
		fmt.Printf("typeName: %v\n", *typName)
	}

	const genSuffix = "_gen.go"

	// figure out reasonable file name if none provided
	if *fileName == "" {
		*fileName = strutil.CamelToSnake(*typName) + genSuffix
	}
	if *verbose {
		fmt.Printf("file name: %v\n", *fileName)
	}

	openFlags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC

	// assume it's OK to overwrite if file was previously generated
	if !strings.HasSuffix(*fileName, genSuffix) {
		openFlags |= os.O_EXCL
	}

	outputFi, err := os.OpenFile(*fileName, openFlags, 0640)
	failErr(errors.Wrapf(err, "failed to open output file %v", *fileName))
	defer func() {
		failErr(errors.Wrap(outputFi.Close(), "failed to close output file"))
	}()

	failErr(errors.Wrap(
		mapgen.Generate(mapgen.Params{
			Exported:   unicode.IsUpper([]rune(*typName)[0]),
			UseRWMutex: *useRwMutex,
			Package:    *pkgName,
			MapName:    *typName,
			KeyType:    key,
			ValType:    value,
		}, outputFi),
		"failed to generate",
	),
	)
	fmt.Printf("Wrote %v\n", outputFi.Name())
	failErr(errors.Wrap(
		exec.Command("goimports", "-w", outputFi.Name()).Run(), "failed to run `goimports`",
	))
}
