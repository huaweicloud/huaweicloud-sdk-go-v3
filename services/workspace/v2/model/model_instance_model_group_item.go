package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceModelGroupItem Agent 实例关联的模型分组信息
type InstanceModelGroupItem struct {

	// 模型分组 ID
	GroupId *string `json:"group_id,omitempty"`

	// 模型分组名称
	GroupName *string `json:"group_name,omitempty"`

	// 分组内默认模型 ID，未设置时为空
	DefaultModelId *string `json:"default_model_id,omitempty"`

	// 分组优先级（数值越小优先级越高）
	Priority *int32 `json:"priority,omitempty"`

	// 模型分组更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 供应商配置列表（不含 API Key）
	Providers *[]InstanceModelProviderConfig `json:"providers,omitempty"`
}

func (o InstanceModelGroupItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceModelGroupItem struct{}"
	}

	return strings.Join([]string{"InstanceModelGroupItem", string(data)}, " ")
}
