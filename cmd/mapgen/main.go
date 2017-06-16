package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/go/loader"

	"strings"

	"os/user"

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

func main() {
	// flags
	var (
		// pkgName       = kingpin.Flag("pkg", "package name").Short('p').Default("main").String()
		verbose       = kingpin.Flag("verbose", "highly descriptive output").Short('v').Bool()
		typName       = kingpin.Flag("tname", "name of generated type").Short('t').String()
		fileName      = kingpin.Flag("fname", "file name of generated type").Short('f').String()
		keyValueTypes = kingpin.Arg("keyvalue type", "Key value type, e.g `string/int`").Required().String()
		loaderArgs    = kingpin.Arg("loader path", "either a series of files or import paths to look for types").Default(".").Strings()
	)
	kingpin.Parse()
	var err error

	// prepare type loader
	if (*loaderArgs)[0] == "." {
		cwd, err := os.Getwd()
		failErr(errors.Wrap(err, "failed to get cwd"))
		(*loaderArgs)[0] = strings.TrimPrefix(cwd, gopath()+"src/")
		if *verbose {
			fmt.Printf("using %v as search path", (*loaderArgs)[0])
		}
	}

	var conf loader.Config
	_, err = conf.FromArgs(*loaderArgs, false)
	if err != nil {
		log.Fatal(err)
	}
	lprog, err := conf.Load()
	failErr(errors.Wrap(err, "failed to load config"))

	pkgs := lprog.InitialPackages()
	if len(pkgs) != 1 {
		log.Fatalf("expected 1 package, got %v", len(pkgs))
	}
	pkg := pkgs[0].Pkg

	// parse basic map info
	key, value, err := mapgen.ParseKeyValueType(*keyValueTypes)
	failErr(err)
	if *verbose {
		fmt.Printf("pkg: %v, key: %v, value: %v\n", pkg.Name(), key, value)
	}

	// figure out reasonable type name
	if *typName == "" {
		*typName = strings.Title(key) + strings.Title(value) + "Map"
	}
	if *verbose {
		fmt.Printf("typeName: %v\n", *typName)
	}

	keyType := pkg.Scope().Lookup(key)
	if keyType == nil {
		fail("failed to find key type %v", key)
	}
	// figure out reasonable file name if none provided
	if *fileName == "" {
		*fileName = strutil.CamelToSnake(*typName) + ".go"
	}
	if *verbose {
		fmt.Printf("file name: %v\n", *fileName)
	}

	outputFi, err := os.OpenFile(*fileName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0640)
	failErr(errors.Wrapf(err, "failed to open output file %v", *fileName))

	failErr(errors.Wrap(outputFi.Close(), "failed to close output file"))
}
