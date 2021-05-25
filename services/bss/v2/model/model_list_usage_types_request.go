package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListUsageTypesRequest struct {
	// 语言。中文：zh_CN英文：en_US缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`
	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。

	ResourceTypeCode *string `json:"resource_type_code,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的数量，默认值为10。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListUsageTypesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListUsageTypesRequest struct{}"
	}

	return strings.Join([]string{"ListUsageTypesRequest", string(data)}, " ")
}
