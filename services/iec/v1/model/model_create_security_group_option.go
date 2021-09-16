package model

import (
	"encoding/json"

	"strings"
)

// 创建安全组参数。
type CreateSecurityGroupOption struct {
	// 安全组的名称。

	Name string `json:"name"`
	// 安全组的描述。非必填项，默认值为空。

	Description *string `json:"description,omitempty"`
}

func (o CreateSecurityGroupOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupOption struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupOption", string(data)}, " ")
}
