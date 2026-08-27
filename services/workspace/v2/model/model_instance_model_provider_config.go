package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceModelProviderConfig 供应商配置（不含 API Key）
type InstanceModelProviderConfig struct {

	// 供应商配置主键 ID
	Id *string `json:"id,omitempty"`

	// 供应商标识
	ProviderId *string `json:"provider_id,omitempty"`

	// 供应商名称
	Name *string `json:"name,omitempty"`

	// 供应商类型
	ProviderType *string `json:"provider_type,omitempty"`

	// 供应商更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 供应商 API 地址
	ApiBaseUrl *string `json:"api_base_url,omitempty"`

	CustomConfig *ProviderCustomConfig `json:"custom_config,omitempty"`

	// 模型列表
	Models *[]ModelInfo `json:"models,omitempty"`
}

func (o InstanceModelProviderConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceModelProviderConfig struct{}"
	}

	return strings.Join([]string{"InstanceModelProviderConfig", string(data)}, " ")
}
