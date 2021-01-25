package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListAuthProjectsRequest struct {
}

func (o KeystoneListAuthProjectsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListAuthProjectsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListAuthProjectsRequest", string(data)}, " ")
}
