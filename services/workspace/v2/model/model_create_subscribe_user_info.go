package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscribeUserInfo 订阅用户信息
type CreateSubscribeUserInfo struct {

	// 用户id。
	UserId *string `json:"user_id,omitempty"`
}

func (o CreateSubscribeUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscribeUserInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscribeUserInfo", string(data)}, " ")
}
