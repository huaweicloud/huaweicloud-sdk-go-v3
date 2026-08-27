package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProviderReq 新增供应商配置请求。
type CreateProviderReq struct {

	// 供应商类型（模板创建时与模板保持一致，自定义时为custom）。
	ProviderType string `json:"provider_type"`

	// 供应商标识（模板创建时与模板保持一致，自定义时可指定）。
	ProviderId string `json:"provider_id"`

	// 供应商API Key（SCC加密存储）。
	ApiKey *string `json:"api_key,omitempty"`

	// 供应商名称（租户自定义）。
	ProviderName string `json:"provider_name"`

	// 自定义Base URL。
	BaseUrl string `json:"base_url"`

	CustomConfig *ProviderCustomConfig `json:"custom_config,omitempty"`

	// 批量创建关联的模型列表。
	Models *[]CreateModelReq `json:"models,omitempty"`

	ApiType *ApiType `json:"api_type"`
}

func (o CreateProviderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProviderReq struct{}"
	}

	return strings.Join([]string{"CreateProviderReq", string(data)}, " ")
}
