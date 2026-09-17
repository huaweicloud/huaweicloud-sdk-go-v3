package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSchemaResponse schema_resource_schema_response
type ResourceSchemaResponse struct {

	// 资源类型
	Type *string `json:"type,omitempty"`

	// schema 内容
	Schema map[string]interface{} `json:"schema,omitempty"`
}

func (o ResourceSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSchemaResponse struct{}"
	}

	return strings.Join([]string{"ResourceSchemaResponse", string(data)}, " ")
}
