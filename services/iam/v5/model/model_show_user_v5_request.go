package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserV5Request Request Object
type ShowUserV5Request struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`
}

func (o ShowUserV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserV5Request struct{}"
	}

	return strings.Join([]string{"ShowUserV5Request", string(data)}, " ")
}
