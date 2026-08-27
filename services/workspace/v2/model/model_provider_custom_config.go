package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProviderCustomConfig 供应商自定义配置，用于指定模型列表接口和认证参数。
type ProviderCustomConfig struct {

	// 模型列表接口地址，用于查询供应商远程模型。
	ModelListApi *string `json:"model_list_api,omitempty"`

	// 认证请求头名称。
	AuthHeader *string `json:"auth_header,omitempty"`

	// 认证前缀（如Bearer）。
	AuthPrefix *string `json:"auth_prefix,omitempty"`

	// 供应商模型列表中模型ID字段名。
	ModelIdField *string `json:"model_id_field,omitempty"`

	// 供应商模型列表中模型名称字段名。
	ModelNameField *string `json:"model_name_field,omitempty"`

	// 自定义HTTP请求头，调用供应商API时附加。
	Headers map[string]string `json:"headers,omitempty"`
}

func (o ProviderCustomConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProviderCustomConfig struct{}"
	}

	return strings.Join([]string{"ProviderCustomConfig", string(data)}, " ")
}
