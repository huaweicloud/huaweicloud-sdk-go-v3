package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductRspBasic 基础版商品编码对象
type ProductRspBasic struct {

	// 云服务产品的主服务类型，云脑默认为：xxx.service.type.sa
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型编码
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源规格编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 资源容量度量标识
	ResourceSizeMeasureId *int32 `json:"resource_size_measure_id,omitempty"`

	// 使用量因子，按需计费必填，取值和话单中的使用量因子一致，云服务和使用量因子对应关系如下: 云脑目前支持有： duration： 时间，主要针对主版本(basic、standard、professional) count：次数，主要针对安全编排 flow：流量，主要针对日志分析和采集 retention：保留，主要针对日志保留
	UsageFactor *string `json:"usage_factor,omitempty"`

	// 使用量单位标识，按需询价必填，例如按小时询价，使用量值为1，使用量单位为小时，枚举值如下： 4：小时 10：GB（带宽按流量询价使用） 11：MB（带宽按流量询价使用） 13：Byte（带宽按流量询价使用）
	UsageMeasureId *int32 `json:"usage_measure_id,omitempty"`

	// 当前region编码，默认为null，即为当前region
	RegionId *string `json:"region_id,omitempty"`
}

func (o ProductRspBasic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductRspBasic struct{}"
	}

	return strings.Join([]string{"ProductRspBasic", string(data)}, " ")
}
