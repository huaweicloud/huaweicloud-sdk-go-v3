package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductInfo 产品信息。
type ProductInfo struct {

	// ID标识，同一次询价中不能重复，用于标识返回询价结果和请求的映射关系。
	Id *string `json:"id,omitempty"`

	// 用户购买云服务产品的云服务类型，例如EC2，云服务类型为hws.service.type.ec2。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 用户购买云服务产品的资源类型，例如EC2中的VM，资源类型为hws.resource.type.vm。
	ResourceType *string `json:"resource_type,omitempty"`

	// 用户购买云服务产品的资源规格，例如VM的小型规格，资源规格为m1.tiny。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 资源容量度量标识。
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// 使用量因子，按需计费必填，取值和话单中的使用量因子一致，云服务和使用量因子对应关系如下： - Duration：云服务器 - flow：流量
	UsageFactor *string `json:"usage_factor,omitempty"`

	// 使用量值，按需询价必填，例如按小时询价，使用量值为1，使用量单位为小时。
	UsageValue *float64 `json:"usage_value,omitempty"`

	// 使用量单位标识，按需询价必填，例如按小时询价，使用量值为1，使用量单位为小时。 - 4：小时 - 10：GB - 11：MB - 13：Byte - 17：FLOW_GB
	UsageMeasureId *int32 `json:"usage_measure_id,omitempty"`

	// 资源容量大小，例如购买的卷大小或带宽大小。
	ResourceSizeMeasureId *int32 `json:"resource_size_measure_id,omitempty"`
}

func (o ProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfo struct{}"
	}

	return strings.Join([]string{"ProductInfo", string(data)}, " ")
}
