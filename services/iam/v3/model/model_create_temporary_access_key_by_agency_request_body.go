package model

import (
	"encoding/json"

	"strings"
)

//
type CreateTemporaryAccessKeyByAgencyRequestBody struct {
	Auth *AgencyAuth `json:"auth"`
}

func (o CreateTemporaryAccessKeyByAgencyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemporaryAccessKeyByAgencyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTemporaryAccessKeyByAgencyRequestBody", string(data)}, " ")
}
