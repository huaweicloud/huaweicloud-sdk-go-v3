package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionUsedInfo struct {

	// 关联资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 管理租户账号
	Owner *string `json:"owner,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`
}

func (o SubscriptionUsedInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionUsedInfo struct{}"
	}

	return strings.Join([]string{"SubscriptionUsedInfo", string(data)}, " ")
}
