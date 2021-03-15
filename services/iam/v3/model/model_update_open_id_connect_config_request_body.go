package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateOpenIdConnectConfigRequestBody struct {
	OpenidConnectConfig *UpdateOpenIdConnectConfig `json:"openid_connect_config"`
}

func (o UpdateOpenIdConnectConfigRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateOpenIdConnectConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateOpenIdConnectConfigRequestBody", string(data)}, " ")
}
