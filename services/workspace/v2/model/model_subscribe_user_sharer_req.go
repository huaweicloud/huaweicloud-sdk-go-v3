package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeUserSharerReq 用户订阅协同请求体。
type SubscribeUserSharerReq struct {

	// 用户协同资源SKU码。
	UserSharerSku string `json:"user_sharer_sku"`

	// 开通协同的的用户列表。
	Users []SubscribeUserInfo `json:"users"`
}

func (o SubscribeUserSharerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeUserSharerReq struct{}"
	}

	return strings.Join([]string{"SubscribeUserSharerReq", string(data)}, " ")
}
