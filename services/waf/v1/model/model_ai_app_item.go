package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiAppItem struct {

	// AI应用ID
	Id *string `json:"id,omitempty"`

	// AI应用名称
	Name *string `json:"name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// API签名
	ApiSign *string `json:"api_sign,omitempty"`

	// 防护状态（0：禁用，1：启用）
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	Policy *Policy `json:"policy,omitempty"`

	// 创建时间戳（毫秒Unix时间戳）
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间戳（毫秒Unix时间戳）
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o AiAppItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiAppItem struct{}"
	}

	return strings.Join([]string{"AiAppItem", string(data)}, " ")
}
