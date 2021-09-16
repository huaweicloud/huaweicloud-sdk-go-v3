package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePublicIpResponse struct {
	Publicip       *PublicIp `json:"publicip,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdatePublicIpResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePublicIpResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicIpResponse", string(data)}, " ")
}
