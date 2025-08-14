package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceResponseSchemaConfigurationReqBody UpdateApplicationInstanceResponseSchemaConfiguration的请求体
type UpdateApplicationInstanceResponseSchemaConfigurationReqBody struct {
	ResponseSchemaConfig *ResponseSchemaConfigDto `json:"response_schema_config"`
}

func (o UpdateApplicationInstanceResponseSchemaConfigurationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceResponseSchemaConfigurationReqBody struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceResponseSchemaConfigurationReqBody", string(data)}, " ")
}
