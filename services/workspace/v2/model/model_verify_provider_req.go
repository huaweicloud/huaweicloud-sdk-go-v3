package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyProviderReq 验证供应商配置请求。
type VerifyProviderReq struct {

	// 供应商主键ID。传入时，其他空字段从数据库已保存的供应商记录中补充。
	Id *string `json:"id,omitempty"`

	// 供应商类型。
	ProviderType *string `json:"provider_type,omitempty"`

	// 供应商id（从模板实例化后的ID）。
	ProviderId *string `json:"provider_id,omitempty"`

	// 供应商API Key（SCC加密存储）。
	ApiKey *string `json:"api_key,omitempty"`

	ApiType *ApiType `json:"api_type,omitempty"`

	// 供应商base_url。
	BaseUrl *string `json:"base_url,omitempty"`

	CustomConfig *ProviderCustomConfig `json:"custom_config,omitempty"`

	// 用于验证连接的模型ID。调用Chat Completion接口时作为model参数传入。
	ModelId *string `json:"model_id,omitempty"`
}

func (o VerifyProviderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyProviderReq struct{}"
	}

	return strings.Join([]string{"VerifyProviderReq", string(data)}, " ")
}
