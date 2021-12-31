package repo

import (
	"errors"
	"fmt"
	"strings"
)

var (
	invalidArguments = errors.New("invalid arguments")
)

func IsBetweenQuery(query *string, arg []interface{}) error {
	*query = strings.TrimSpace(*query)
	if strings.Contains(*query, "?") ||
		strings.Contains(*query, " ") ||
		strings.Contains(*query, "@") ||
		len(arg) != 2 {
		return invalidArguments
	}
	*query = fmt.Sprintf("%s BETWEEN ? AND ?", *query)

	return nil
}
