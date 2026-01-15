package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeUserInfo 用户信息。
type SubscribeUserInfo struct {

	// 用户id。
	UserId string `json:"user_id"`

	// 用户名称。
	UserName string `json:"user_name"`
}

func (o SubscribeUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeUserInfo struct{}"
	}

	return strings.Join([]string{"SubscribeUserInfo", string(data)}, " ")
}
