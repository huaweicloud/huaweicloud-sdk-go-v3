package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreatePublicIpResponse struct {
	Publicip       *PublicIp `json:"publicip,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreatePublicIpResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePublicIpResponse struct{}"
	}

	return strings.Join([]string{"CreatePublicIpResponse", string(data)}, " ")
}
