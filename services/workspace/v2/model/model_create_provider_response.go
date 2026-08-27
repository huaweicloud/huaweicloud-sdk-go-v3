package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProviderResponse Response Object
type CreateProviderResponse struct {

	// 供应商id。
	Id *string `json:"id,omitempty"`

	// 供应商类型。
	ProviderType *string `json:"provider_type,omitempty"`

	// 供应商id。
	ProviderId *string `json:"provider_id,omitempty"`

	// 供应商名称。
	ProviderName *string `json:"provider_name,omitempty"`

	// 供应商base_url。
	BaseUrl *string `json:"base_url,omitempty"`

	// 连接状态（connected/disconnected/unverified）。
	ConnectionStatus *string `json:"connection_status,omitempty"`

	// 是否内置供应商。
	IsBuiltin *bool `json:"is_builtin,omitempty"`

	ApiType *ApiType `json:"api_type,omitempty"`

	// 自定义配置。
	CustomConfig *interface{} `json:"custom_config,omitempty"`

	// 最后验证时间（ISO8601格式，UTC时区）。
	LastVerifyTime *string `json:"last_verify_time,omitempty"`

	// 创建时间（ISO8601格式，UTC时区）。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间（ISO8601格式，UTC时区）。
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建的模型列表。
	CreatedModels  *[]ModelItemResp `json:"created_models,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProviderResponse struct{}"
	}

	return strings.Join([]string{"CreateProviderResponse", string(data)}, " ")
}
