package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModelGroupResp 模型组详情响应，包含关联的供应商简化信息。
type ShowModelGroupResp struct {

	// 分组id。
	Id *string `json:"id,omitempty"`

	// 分组名称。
	Name *string `json:"name,omitempty"`

	// 分组描述。
	Description *string `json:"description,omitempty"`

	// 分组优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 默认模型ID。
	DefaultModelId *string `json:"default_model_id,omitempty"`

	// 关联的供应商简化信息列表（仅id和name）。
	Providers *[]ModelGroupProviderSimpleResp `json:"providers,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间（ISO8601格式，UTC时区）。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ShowModelGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModelGroupResp struct{}"
	}

	return strings.Join([]string{"ShowModelGroupResp", string(data)}, " ")
}
