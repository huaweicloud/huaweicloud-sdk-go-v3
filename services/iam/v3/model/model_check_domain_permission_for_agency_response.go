package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CheckDomainPermissionForAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckDomainPermissionForAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckDomainPermissionForAgencyResponse struct{}"
	}

	return strings.Join([]string{"CheckDomainPermissionForAgencyResponse", string(data)}, " ")
}
