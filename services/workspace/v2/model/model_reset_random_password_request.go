package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetRandomPasswordRequest Request Object
type ResetRandomPasswordRequest struct {

	// 用户ID。
	UserId string `json:"user_id"`

	// 通知用户类型，现在有“email”和“phone”两种，用\",\"分开，用户为用户激活模式时必须要带上通知类型，以便接收到密码通知。
	NotificationType *string `json:"notification_type,omitempty"`
}

func (o ResetRandomPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetRandomPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetRandomPasswordRequest", string(data)}, " ")
}
