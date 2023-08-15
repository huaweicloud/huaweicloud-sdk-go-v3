package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductInfoBean 产品信息
type ProductInfoBean struct {

	// 资源名称列表
	AllResourceNames *[]string `json:"all_resource_names,omitempty"`

	// 云服务类型
	CloudServiceType string `json:"cloud_service_type"`

	// 展示ID
	DisplayId *string `json:"display_id,omitempty"`

	// 产品ID
	ProductId string `json:"product_id"`

	// 产品规格描述
	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 产品支持的数据库数量，或者支持obs的扫描量
	ResourceSize int32 `json:"resource_size"`

	// 资源容量度量标识，枚举值举例如下：15：mbps（购买带宽时使用），17：gb（购买云硬盘时使用），14：个/次
	ResourceSizeMeasureId int32 `json:"resource_size_measure_id"`

	// 产品编码
	ResourceSpecCode string `json:"resource_spec_code"`

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 已使用系数
	UsageFactor *string `json:"usage_factor,omitempty"`

	// 已使用容量度量标识
	UsageMeasureId *int32 `json:"usage_measure_id,omitempty"`

	// 已使用值
	UsageValue *float64 `json:"usage_value,omitempty"`
}

func (o ProductInfoBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfoBean struct{}"
	}

	return strings.Join([]string{"ProductInfoBean", string(data)}, " ")
}
