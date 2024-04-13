package gosteps

import (
	"fmt"
	"strconv"
)

func intParser(raw string) (interface{}, error) {
	rawInt, err := strconv.Atoi(raw)
	if err != nil {
		return nil, err
	}
	return rawInt, nil
}

func intComparer(raw string, actual interface{}) error {
	as, ok := actual.(int)
	if !ok {
		return fmt.Errorf("actual value noot int")
	}

	rawInt, err := strconv.Atoi(raw)
	if err != nil {
		return err
	}
	if as != rawInt {
		return fmt.Errorf("value mismatch")
	}
	return nil
}
