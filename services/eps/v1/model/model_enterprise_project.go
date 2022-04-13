package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 企业项目
type EnterpriseProject struct {
	// 只能由中文字符、英文字母（a~zA~Z）、数字（0~9）、下划线（_）、中划线（-）组成，且长度为[1-64]个字符。名称不能为大小写混合的default，且在租户账号内唯一。

	Name string `json:"name"`
	// 最大长度512个字符。

	Description *string `json:"description,omitempty"`
	// 企业项目类型

	Type *string `json:"type,omitempty"`
}

func (o EnterpriseProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProject struct{}"
	}

	return strings.Join([]string{"EnterpriseProject", string(data)}, " ")
}
