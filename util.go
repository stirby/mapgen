package mapgen

import (
	"strings"

	"github.com/pkg/errors"
)

// ParseKeyValueType parses a string of the form <key>/<value>
func ParseKeyValueType(kvt string) (key string, value string, err error) {
	tokens := strings.Split(kvt, "/")
	if len(tokens) != 2 {
		return "", "", errors.Errorf("unexpected token count %v", len(tokens))
	}
	return tokens[0], tokens[1], nil
}
