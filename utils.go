package zorro

import (
	"errors"
	"fmt"
	"strings"
)

func JoinErrors(errs ...error) error {
	var parts []string
	for _, err := range errs {
		if err != nil {
			parts = append(parts, err.Error())
		}
	}
	if len(parts) == 0 {
		return nil
	}
	return errors.New(fmt.Sprintf("errors: %s", strings.Join(parts, ", ")))
}
