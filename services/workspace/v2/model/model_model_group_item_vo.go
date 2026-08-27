package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelGroupItemVo 模型组列表项响应（不包含providers）。
type ModelGroupItemVo struct {

	// 模型组id。
	Id *string `json:"id,omitempty"`

	// 分组名称。
	Name *string `json:"name,omitempty"`

	// 分组描述。
	Description *string `json:"description,omitempty"`

	// 分组优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 默认模型ID。
	DefaultModelId *string `json:"default_model_id,omitempty"`

	// 关联的供应商数量。
	ProviderCount *int32 `json:"provider_count,omitempty"`

	// 关联的应用对象数量（Agent实例+桌面标签）。
	ResourceCount *int32 `json:"resource_count,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间（ISO8601格式，UTC时区）。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ModelGroupItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelGroupItemVo struct{}"
	}

	return strings.Join([]string{"ModelGroupItemVo", string(data)}, " ")
}
