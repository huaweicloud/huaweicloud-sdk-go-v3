package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type DemandProductInfo struct {

	// ID标识，同一次询价中不能重复，用于标识返回询价结果和请求的映射关系。
	Id string `json:"id"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。
	CloudServiceType string `json:"cloud_service_type"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。 ResourceType是CloudServiceType中的一种资源，CloudServiceType由多种ResourceType组合提供。
	ResourceType string `json:"resource_type"`

	// 云服务类型的资源规格，部分云服务类型和资源规格举例如下： 弹性云服务器：根据操作系统类型在云服务器规格的ID后添加“.win”或“.linux”，例如“s2.small.1.linux”。云服务器规格的ID字段，您可以调用查询规格详情和规格扩展信息列表接口获取。 带宽：12_bgp：动态BGP按流量计费带宽12_sbgp：静态BGP按流量计费带宽19_bgp：动态BGP按带宽计费带宽19_sbgp：静态BGP按带宽计费带宽19_share：按带宽计费共享带宽 IP：5_bgp：动态BGP公网IP5_sbgp：静态BGP公网IP 云硬盘：SATA：普通IO云硬盘SAS：高IO云硬盘GPSSD：通用型SSD云硬盘SSD：超高IO云硬盘
	ResourceSpec string `json:"resource_spec"`

	// 云服务区编码，例如：“cn-north-1”。
	Region string `json:"region"`

	// 可用区标识。此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	AvailableZone *string `json:"available_zone,omitempty"`

	// 资源容量大小，例如购买的卷大小或带宽大小。 线性产品时该参数不能为空。线性产品为包括硬盘，带宽等在订购时需要指定大小的产品。例如硬盘在订购时需选择10G、20G等不同大小。非线性产品时此参数不携带或者携带值为null时，不作为筛选条件。
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// 资源容量度量标识，枚举值如下： 15：Mbps（购买带宽时使用）17：GB（购买云硬盘时使用）14：个（次） 线性产品时该参数不能为空。线性产品为包括硬盘，带宽等在订购时需要指定大小的产品。例如硬盘在订购时需选择10G、20G等不同大小。非线性产品时此参数不携带或者携带值为null时，不作为筛选条件。
	SizeMeasureId *int32 `json:"size_measure_id,omitempty"`

	// 使用量因子编码，大小写不敏感，取值和话单中的使用量因子一致，云服务和使用量因子对应关系举例如下： 云服务器：Duration云硬盘：Duration弹性IP：Duration带宽：Duration或upflow市场镜像：Duration 您可以调用查询使用量类型列表接口获取响应参数表3中参数code的取值，即每种云服务对应的计费因子。
	UsageFactor string `json:"usage_factor"`

	// 使用量值。 例如按小时询价，使用量值为1，使用量单位为小时。
	UsageValue *decimal.Decimal `json:"usage_value"`

	// 使用量度量单位， 例如按小时询价，使用量值为1，使用量单位为小时。
	UsageMeasureId int32 `json:"usage_measure_id"`

	// 订购数量。
	SubscriptionNum int32 `json:"subscription_num"`
}

func (o DemandProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DemandProductInfo struct{}"
	}

	return strings.Join([]string{"DemandProductInfo", string(data)}, " ")
}
