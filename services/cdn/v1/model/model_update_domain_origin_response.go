package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDomainOriginResponse struct {
	Origin         *ResourceBody `json:"origin,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateDomainOriginResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainOriginResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainOriginResponse", string(data)}, " ")
}
