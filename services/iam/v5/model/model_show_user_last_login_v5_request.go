package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserLastLoginV5Request Request Object
type ShowUserLastLoginV5Request struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`
}

func (o ShowUserLastLoginV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserLastLoginV5Request struct{}"
	}

	return strings.Join([]string{"ShowUserLastLoginV5Request", string(data)}, " ")
}
