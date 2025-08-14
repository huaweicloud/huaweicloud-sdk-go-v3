package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResponseSchemaConfigDto 应用程序Schema属性映射配置
type ResponseSchemaConfigDto struct {

	// 额外添加的属性映射Schema配置
	Properties map[string]ResponseSchemaPropertiesDetailsDto `json:"properties,omitempty"`

	Subject *ResponseSchemaSubjectDetailsDto `json:"subject"`

	// 应用程序支持的Subject NameID格式
	SupportedNameIdFormats *[]string `json:"supported_name_id_formats,omitempty"`
}

func (o ResponseSchemaConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseSchemaConfigDto struct{}"
	}

	return strings.Join([]string{"ResponseSchemaConfigDto", string(data)}, " ")
}
