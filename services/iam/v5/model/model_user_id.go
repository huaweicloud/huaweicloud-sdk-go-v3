package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserId IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
type UserId struct {
}

func (o UserId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserId struct{}"
	}

	return strings.Join([]string{"UserId", string(data)}, " ")
}
