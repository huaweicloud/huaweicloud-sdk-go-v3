package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceSpecInfo struct {

	// 云服务类型编码
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 云服务类型名称
	CloudServiceTypeName *string `json:"cloud_service_type_name,omitempty"`

	// 资源类型编码
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源类型名称
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 云服务类型的资源规格编码
	ResourceSpec *string `json:"resource_spec,omitempty"`

	// 云服务类型的资源规格名称
	ResourceSpecName *string `json:"resource_spec_name,omitempty"`

	// 属性列表，need_attributes=true时返回属性信息。
	Attributes *[]Attribute `json:"attributes,omitempty"`

	// 定价列表，need_price=true时返回定价信息。
	PriceLists *[]PriceItem `json:"price_lists,omitempty"`
}

func (o ResourceSpecInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpecInfo struct{}"
	}

	return strings.Join([]string{"ResourceSpecInfo", string(data)}, " ")
}
