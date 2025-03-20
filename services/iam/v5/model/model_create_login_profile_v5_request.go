package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoginProfileV5Request Request Object
type CreateLoginProfileV5Request struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`

	Body *CreateLoginProfileReqBody `json:"body,omitempty"`
}

func (o CreateLoginProfileV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoginProfileV5Request struct{}"
	}

	return strings.Join([]string{"CreateLoginProfileV5Request", string(data)}, " ")
}
