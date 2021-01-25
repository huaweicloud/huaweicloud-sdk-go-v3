package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListIdentityProvidersRequest struct {
}

func (o KeystoneListIdentityProvidersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListIdentityProvidersRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListIdentityProvidersRequest", string(data)}, " ")
}
