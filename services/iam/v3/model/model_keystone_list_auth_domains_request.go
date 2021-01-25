package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListAuthDomainsRequest struct {
}

func (o KeystoneListAuthDomainsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListAuthDomainsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListAuthDomainsRequest", string(data)}, " ")
}
