package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserV5Request Request Object
type DeleteUserV5Request struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`
}

func (o DeleteUserV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserV5Request struct{}"
	}

	return strings.Join([]string{"DeleteUserV5Request", string(data)}, " ")
}
