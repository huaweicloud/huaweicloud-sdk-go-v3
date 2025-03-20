package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserV5Request Request Object
type UpdateUserV5Request struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`

	Body *UpdateUserReqBody `json:"body,omitempty"`
}

func (o UpdateUserV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserV5Request struct{}"
	}

	return strings.Join([]string{"UpdateUserV5Request", string(data)}, " ")
}
