package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPublicIpResponse struct {
	Publicip       *PublicIp `json:"publicip,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowPublicIpResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPublicIpResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicIpResponse", string(data)}, " ")
}
