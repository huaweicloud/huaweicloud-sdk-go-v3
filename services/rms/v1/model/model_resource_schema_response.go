package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单条schema信息
type ResourceSchemaResponse struct {

	// 资源类型.
	Type *string `json:"type,omitempty" xml:"type"`

	// schema 内容.
	Schema *interface{} `json:"schema,omitempty" xml:"schema"`
}

func (o ResourceSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSchemaResponse struct{}"
	}

	return strings.Join([]string{"ResourceSchemaResponse", string(data)}, " ")
}
