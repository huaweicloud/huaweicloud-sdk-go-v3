package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInfoBeanNew struct {

	// 所有资源名称
	AllResourceNames *[]string `json:"all_resource_names,omitempty"`

	// 云服务类型
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 显示ID
	DisplayId *string `json:"display_id,omitempty"`

	// 产品ID
	ProductId *string `json:"product_id,omitempty"`

	// 产品规格描述
	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源规格
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 资源类型: - hws.resource.type.dbss：数据库审计
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o ProductInfoBeanNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfoBeanNew struct{}"
	}

	return strings.Join([]string{"ProductInfoBeanNew", string(data)}, " ")
}
