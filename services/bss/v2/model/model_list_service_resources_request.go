package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListServiceResourcesRequest struct {
	// 语言。中文：zh_CN英文：en_US缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`
	// 云服务类型编码。例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。

	ServiceTypeCode string `json:"service_type_code"`
	// 每次查询的数量，默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
}

func (o ListServiceResourcesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListServiceResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListServiceResourcesRequest", string(data)}, " ")
}
