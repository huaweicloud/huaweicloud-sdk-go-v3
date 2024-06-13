package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIfUserNameRepeatRequest Request Object
type ShowIfUserNameRepeatRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// 用户ID
	UserId *string `json:"userId,omitempty"`

	// 用户名
	UserName *string `json:"userName,omitempty"`
}

func (o ShowIfUserNameRepeatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIfUserNameRepeatRequest struct{}"
	}

	return strings.Join([]string{"ShowIfUserNameRepeatRequest", string(data)}, " ")
}
