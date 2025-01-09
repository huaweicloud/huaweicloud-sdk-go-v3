package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeUserSharerReq 用户订阅协同请求体
type SubscribeUserSharerReq struct {

	// 用户协同资源SKU码
	UserSharerSku string `json:"user_sharer_sku"`

	// 订单ID
	OrderId *string `json:"order_id,omitempty"`

	// 开通协同的的用户列表。
	Users []User `json:"users"`

	// 企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o SubscribeUserSharerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeUserSharerReq struct{}"
	}

	return strings.Join([]string{"SubscribeUserSharerReq", string(data)}, " ")
}
