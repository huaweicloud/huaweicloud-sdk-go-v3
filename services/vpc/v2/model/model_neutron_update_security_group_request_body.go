package model

import (
	"encoding/json"

	"strings"
)

//
type NeutronUpdateSecurityGroupRequestBody struct {
	SecurityGroup *NeutronUpdateSecurityGroupOption `json:"security_group"`
}

func (o NeutronUpdateSecurityGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronUpdateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSecurityGroupRequestBody", string(data)}, " ")
}
