package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLoginProfileV5Request Request Object
type UpdateLoginProfileV5Request struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`

	Body *UpdateLoginProfileReqBody `json:"body,omitempty"`
}

func (o UpdateLoginProfileV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoginProfileV5Request struct{}"
	}

	return strings.Join([]string{"UpdateLoginProfileV5Request", string(data)}, " ")
}
