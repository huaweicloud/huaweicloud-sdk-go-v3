package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserName IAM用户名，长度为1到64个字符，只包含字母、数字、\"_\"、\"-\"、\".\"和空格的字符串，且首位不能为数字。
type UserName struct {
}

func (o UserName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserName struct{}"
	}

	return strings.Join([]string{"UserName", string(data)}, " ")
}
