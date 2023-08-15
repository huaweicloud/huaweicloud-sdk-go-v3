package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSchemaResponse 单条schema信息
type ResourceSchemaResponse struct {

	// 资源类型.
	Type *string `json:"type,omitempty"`

	// schema 内容.
	Schema *interface{} `json:"schema,omitempty"`
}

func (o ResourceSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSchemaResponse struct{}"
	}

	return strings.Join([]string{"ResourceSchemaResponse", string(data)}, " ")
}
