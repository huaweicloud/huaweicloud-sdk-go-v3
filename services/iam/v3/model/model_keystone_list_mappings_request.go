package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListMappingsRequest struct {
}

func (o KeystoneListMappingsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListMappingsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListMappingsRequest", string(data)}, " ")
}
