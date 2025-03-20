package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyName 身份策略名称，长度为1到128个字符，只包含字母、数字、\"_\"、\"+\"、\"=\"、\".\"、\"@\"和\"-\"的字符串。
type PolicyName struct {
}

func (o PolicyName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyName struct{}"
	}

	return strings.Join([]string{"PolicyName", string(data)}, " ")
}
