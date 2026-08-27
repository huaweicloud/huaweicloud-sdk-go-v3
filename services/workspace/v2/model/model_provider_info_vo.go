package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProviderInfoVo 供应商配置信息（含last_verify_time）。
type ProviderInfoVo struct {

	// 供应商id。
	Id *string `json:"id,omitempty"`

	// 供应商名称。
	ProviderName *string `json:"provider_name,omitempty"`

	// 供应商类型。
	ProviderType *string `json:"provider_type,omitempty"`

	// 供应商id。
	ProviderId *string `json:"provider_id,omitempty"`

	// 供应商base_url。
	BaseUrl *string `json:"base_url,omitempty"`

	// 连接状态（connected/disconnected/unverified）。
	ConnectionStatus *string `json:"connection_status,omitempty"`

	// 下属模型数量。
	ModelCount *int32 `json:"model_count,omitempty"`

	// 关联的模型分组数量。
	GroupCount *int32 `json:"group_count,omitempty"`

	// 最后一次验证时间。
	LastVerifyTime *string `json:"last_verify_time,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间（ISO8601格式，UTC时区）。
	UpdateTime *string `json:"update_time,omitempty"`

	// 是否为内置供应商。
	IsBuiltin *bool `json:"is_builtin,omitempty"`

	ApiType *ApiType `json:"api_type,omitempty"`
}

func (o ProviderInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProviderInfoVo struct{}"
	}

	return strings.Join([]string{"ProviderInfoVo", string(data)}, " ")
}
