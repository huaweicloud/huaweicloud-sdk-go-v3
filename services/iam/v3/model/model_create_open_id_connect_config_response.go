package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateOpenIdConnectConfigResponse struct {
	OpenidConnectConfig *CreateOpenIdConnectConfig `json:"openid_connect_config,omitempty"`
	HttpStatusCode      int                        `json:"-"`
}

func (o CreateOpenIdConnectConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateOpenIdConnectConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateOpenIdConnectConfigResponse", string(data)}, " ")
}
