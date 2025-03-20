package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyId 身份策略ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
type PolicyId struct {
}

func (o PolicyId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyId struct{}"
	}

	return strings.Join([]string{"PolicyId", string(data)}, " ")
}
