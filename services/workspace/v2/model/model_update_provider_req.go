package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProviderReq 更新供应商配置请求。
type UpdateProviderReq struct {

	// 供应商标识（模板创建时与模板保持一致，自定义时可指定）。
	ProviderId *string `json:"provider_id,omitempty"`

	// 供应商名称。
	ProviderName *string `json:"provider_name,omitempty"`

	// 供应商base_url。
	BaseUrl *string `json:"base_url,omitempty"`

	// 供应商API Key（SCC加密存储）。
	ApiKey *string `json:"api_key,omitempty"`

	CustomConfig *ProviderCustomConfig `json:"custom_config,omitempty"`

	ApiType *ApiType `json:"api_type,omitempty"`
}

func (o UpdateProviderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProviderReq struct{}"
	}

	return strings.Join([]string{"UpdateProviderReq", string(data)}, " ")
}
