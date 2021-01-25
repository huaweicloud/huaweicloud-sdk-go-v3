package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneCreateAgencyTokenRequestBody struct {
	Auth *AgencyTokenAuth `json:"auth"`
}

func (o KeystoneCreateAgencyTokenRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneCreateAgencyTokenRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateAgencyTokenRequestBody", string(data)}, " ")
}
