package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrderReq 购买参数
type OrderReq struct {

	// 区域Id
	RegionId string `json:"region_id"`

	// 订单Id
	CommodityId *string `json:"commodity_id,omitempty"`

	// 产品Id
	ProductId *string `json:"product_id,omitempty"`

	// 购买周期类型（日月年）
	PeriodType int32 `json:"period_type"`

	// 购买周期数
	PeriodNum int32 `json:"period_num"`

	// 可用区
	AvailabilityZone string `json:"availability_zone"`

	// 虚拟网卡Id
	VpcId string `json:"vpc_id"`

	// 安全组Id
	SecurityGroupId string `json:"security_group_id"`

	// 子网Id
	NetId string `json:"net_id"`

	// 实例名
	InstanceName string `json:"instance_name"`

	// 企业项目Id
	EpsId string `json:"eps_id"`

	// 是否续订
	IsAutoRenew int32 `json:"is_auto_renew"`

	// 促销信息
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 实例附加增量包类型
	ExtesionPackageType *string `json:"extesion_package_type,omitempty"`

	// 增量包绑定的实例id
	BindingInstanceId *string `json:"binding_instance_id,omitempty"`

	// cdm版本号
	CdmVersion *string `json:"cdm_version,omitempty"`

	// 实例规格编码
	ResourceSpecCode string `json:"resource_spec_code"`

	// 云服务类型
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// tms标签
	Tags *[]TmsTagDto `json:"tags,omitempty"`
}

func (o OrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrderReq struct{}"
	}

	return strings.Join([]string{"OrderReq", string(data)}, " ")
}
