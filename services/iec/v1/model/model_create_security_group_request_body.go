package model

import (
	"encoding/json"

	"strings"
)

// 创建安全组请求体。
type CreateSecurityGroupRequestBody struct {
	SecurityGroup *CreateSecurityGroupOption `json:"security_group"`
}

func (o CreateSecurityGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRequestBody", string(data)}, " ")
}
