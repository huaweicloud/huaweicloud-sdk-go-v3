package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneCreateProtocolResponse struct {
	Protocol       *ProtocolResult `json:"protocol,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o KeystoneCreateProtocolResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateProtocolResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProtocolResponse", string(data)}, " ")
}
