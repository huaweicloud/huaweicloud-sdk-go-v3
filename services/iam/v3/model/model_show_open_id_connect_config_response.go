package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowOpenIdConnectConfigResponse struct {
	OpenidConnectConfig *OpenIdConnectConfig `json:"openid_connect_config,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowOpenIdConnectConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowOpenIdConnectConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowOpenIdConnectConfigResponse", string(data)}, " ")
}
