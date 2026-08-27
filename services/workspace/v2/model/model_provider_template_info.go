package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProviderTemplateInfo 供应商模板信息。
type ProviderTemplateInfo struct {

	// 模板唯一标识（供应商类型）。
	ProviderType *string `json:"provider_type,omitempty"`

	// 供应商id。
	ProviderId *string `json:"provider_id,omitempty"`

	// 供应商base_url。
	BaseUrl *string `json:"base_url,omitempty"`

	CustomConfig *ProviderCustomConfig `json:"custom_config,omitempty"`

	ApiType *ApiType `json:"api_type,omitempty"`
}

func (o ProviderTemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProviderTemplateInfo struct{}"
	}

	return strings.Join([]string{"ProviderTemplateInfo", string(data)}, " ")
}
