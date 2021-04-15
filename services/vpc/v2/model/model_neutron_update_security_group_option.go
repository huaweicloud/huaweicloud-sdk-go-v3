package model

import (
	"encoding/json"

	"strings"
)

//
type NeutronUpdateSecurityGroupOption struct {
	// 功能说明：安全组描述 取值范围：0-255个字符

	Description *string `json:"description,omitempty"`
	// 功能说明：安全组名称 取值范围：0-255个字符 约束：不允许为“default”

	Name *string `json:"name,omitempty"`
}

func (o NeutronUpdateSecurityGroupOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronUpdateSecurityGroupOption struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSecurityGroupOption", string(data)}, " ")
}
