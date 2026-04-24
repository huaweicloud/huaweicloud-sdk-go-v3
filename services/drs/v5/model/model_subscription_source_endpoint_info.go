package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionSourceEndpointInfo struct {

	// 数据库实例ID
	Id string `json:"id"`

	// 数据库实例类型，仅支持mysql
	Type *string `json:"type,omitempty"`
}

func (o SubscriptionSourceEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionSourceEndpointInfo struct{}"
	}

	return strings.Join([]string{"SubscriptionSourceEndpointInfo", string(data)}, " ")
}
