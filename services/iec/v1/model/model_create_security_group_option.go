package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupOption struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupOption", string(data)}, " ")
}
