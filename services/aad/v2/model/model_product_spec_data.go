package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductSpecData struct {

	// 产品规格UUID
	ProductUuid *string `json:"product_uuid,omitempty"`

	// 套餐线路
	IspSpec *string `json:"isp_spec,omitempty"`

	// 机房
	DataCenter *string `json:"data_center,omitempty"`

	// 产品规格类型
	SpecType *int32 `json:"spec_type,omitempty"`

	// 保底带宽值
	BasicBandwidth *int32 `json:"basic_bandwidth,omitempty"`

	// 弹性带宽值
	ElasticBandwidth *int32 `json:"elastic_bandwidth,omitempty"`

	// 业务带宽值
	ServiceBandwidth *int32 `json:"service_bandwidth,omitempty"`

	// 端口数
	PortNum *int32 `json:"port_num,omitempty"`

	// 域名数
	BindDomainNum *int32 `json:"bind_domain_num,omitempty"`

	// 弹性业务带宽值
	ElasticServiceBandwidth *int32 `json:"elastic_service_bandwidth,omitempty"`

	// 弹性业务带宽类型
	ElasticServiceBandwidthType *int32 `json:"elastic_service_bandwidth_type,omitempty"`

	// 主资源类型
	MainResourceType *string `json:"main_resource_type,omitempty"`

	// 主资源规格编码
	MainResourceSpecCode *string `json:"main_resource_spec_code,omitempty"`

	// 主资源产品id
	MainResourceProductId *string `json:"main_resource_product_id,omitempty"`
}

func (o ProductSpecData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductSpecData struct{}"
	}

	return strings.Join([]string{"ProductSpecData", string(data)}, " ")
}
