package model

import (
	"encoding/json"

	"strings"
)

//
type NeutronCreateSecurityGroupRequestBody struct {
	SecurityGroup *NeutronCreateSecurityGroupOption `json:"security_group"`
}

func (o NeutronCreateSecurityGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRequestBody", string(data)}, " ")
}
