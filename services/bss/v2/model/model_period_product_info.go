package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeriodProductInfo struct {

	// ID标识，同一次询价中不能重复，用于标识返回询价结果和请求的映射关系。
	Id string `json:"id"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。
	CloudServiceType string `json:"cloud_service_type"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。 ResourceType是CloudServiceType中的一种资源，CloudServiceType由多种ResourceType组合提供。
	ResourceType string `json:"resource_type"`

	// 云服务类型的资源规格，部分云服务类型和资源规格举例如下： 弹性云服务器：根据操作系统类型在云服务器规格的ID后添加“.win”或“.linux”，例如“s2.small.1.linux”。云服务器规格的ID字段，您可以调用查询规格详情和规格扩展信息列表接口获取。 带宽：12_bgp：动态BGP按流量计费带宽12_sbgp：静态BGP按流量计费带宽19_bgp：动态BGP按带宽计费带宽19_sbgp：静态BGP按带宽计费带宽19_share：按带宽计费共享带宽 IP：5_bgp：动态BGP公网IP5_sbgp：静态BGP公网IP 云数据库：云数据库的资源规格信息，您可以调用查询数据库规格接口获取。 分布式缓存服务：分布式缓存服务的资源规格信息，您可以调用查询产品规格接口获取。
	ResourceSpec string `json:"resource_spec"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。
	Region string `json:"region"`

	// 可用区标识，例如：“cn-north-1a”，大小写不敏感。具体请参见地区和终端节点可用分区的“可用分区名称”列的值。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	AvailableZone *string `json:"available_zone,omitempty"`

	// 资源容量大小，例如购买的卷大小或带宽大小。 线性产品时该参数不能为空。线性产品为包括硬盘，带宽等在订购时需要指定大小的产品。例如硬盘在订购时需选择10G、20G等不同大小。非线性产品时此参数不携带或者携带值为null时，不作为筛选条件。
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// 资源容量度量标识。 15：Mbps（购买带宽时使用）17：GB（购买云硬盘时使用）14：个 线性产品时该参数必填。线性产品为包括硬盘，带宽等在订购时需要指定大小的产品。例如硬盘在订购时需选择10G、20G等不同大小。 非线性产品时此参数不携带或者携带值为null时，不作为筛选条件。
	SizeMeasureId *int32 `json:"size_measure_id,omitempty"`

	// 订购包年/包月产品的周期类型。 0：天2：月3：年4：小时
	PeriodType int32 `json:"period_type"`

	// 订购包年/包月产品的周期数。
	PeriodNum int32 `json:"period_num"`

	// 订购包年/包月产品的数量。
	SubscriptionNum int32 `json:"subscription_num"`

	// 费用分期模式。 HALF_PAY：半付ZERO_PAY：零付NA：不支持费用分期模式  说明： 此参数不携带或携带值为空串或携带值为null时，默认值为“NA”。暂只支持CloudPond产品。
	FeeInstallmentMode *string `json:"fee_installment_mode,omitempty"`
}

func (o PeriodProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodProductInfo struct{}"
	}

	return strings.Join([]string{"PeriodProductInfo", string(data)}, " ")
}
