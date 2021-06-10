package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateOriginHostResponse struct {
	OriginHost     *DomainOriginHost `json:"origin_host,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateOriginHostResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateOriginHostResponse struct{}"
	}

	return strings.Join([]string{"UpdateOriginHostResponse", string(data)}, " ")
}
