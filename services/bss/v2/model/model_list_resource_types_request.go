package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListResourceTypesRequest struct {
	// 语言。zh_CN：中文en_US：英文缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`
	// 资源类型编码。例如ECS的VM为“hws.resource.type.vm”。

	ResourceTypeCode *string `json:"resource_type_code,omitempty"`
}

func (o ListResourceTypesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourceTypesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceTypesRequest", string(data)}, " ")
}
