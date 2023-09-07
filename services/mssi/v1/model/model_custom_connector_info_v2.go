package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomConnectorInfoV2 CustomConnector元数据
type CustomConnectorInfoV2 struct {
	ApiConfig *ApiConfig `json:"api_config,omitempty"`

	AuthContent *AuthConfigA `json:"auth_content,omitempty"`

	ConnectorCreatedType *string `json:"connector_created_type,omitempty"`

	ConnectorType *string `json:"connector_type,omitempty"`

	// logo base64编码
	Icon *string `json:"icon,omitempty"`

	// swagger文档，大文本
	Swagger *interface{} `json:"swagger,omitempty"`
}

func (o CustomConnectorInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomConnectorInfoV2 struct{}"
	}

	return strings.Join([]string{"CustomConnectorInfoV2", string(data)}, " ")
}
