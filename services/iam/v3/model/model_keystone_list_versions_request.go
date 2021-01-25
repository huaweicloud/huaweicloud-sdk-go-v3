package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListVersionsRequest struct {
}

func (o KeystoneListVersionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListVersionsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListVersionsRequest", string(data)}, " ")
}
