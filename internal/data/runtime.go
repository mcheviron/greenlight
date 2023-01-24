package data

import (
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = fmt.Errorf("invalid runtime format")

func (runtime *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJsonValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	parts := strings.Split(unquotedJsonValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
	runtimeValue, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*runtime = Runtime(runtimeValue)
	return nil
}

func (r Runtime) MarshalJSON() ([]byte, error) {
	// Add quotation to the string ort it won't be consider as valid JSON
	jsonValue := strconv.Quote(fmt.Sprintf("%d mins", r))
	return []byte(jsonValue), nil
}
