package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChargeOrderDbssNew struct {

	// 资源标签列表
	Tags *[]ResourceTagInfoBean `json:"tags,omitempty"`

	// 资产数量
	AssetNums int32 `json:"asset_nums"`

	// 可用分区
	AvailabilityZone string `json:"availability_zone"`

	// 计费模式
	ChargingMode int32 `json:"charging_mode"`

	// 服务类型
	CloudServiceType string `json:"cloud_service_type"`

	// 备注信息
	Comment *string `json:"comment,omitempty"`

	// 组合套餐
	CompositeProductId *string `json:"composite_product_id,omitempty"`

	// 实例部署方式，默认为云上 - CLOUD：云上 - OUTSIDE：云外
	DeployMode string `json:"deploy_mode"`

	// 折扣ID
	DiscountId *string `json:"discount_id,omitempty"`

	// 企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// ECS规格ID
	FlavorRef string `json:"flavor_ref"`

	// 前端登录密码
	HxPassword string `json:"hx_password"`

	// 镜像业务类型
	ImageBusinessType int32 `json:"image_business_type"`

	// 自费续费  - 1：自动  - 0：不自动
	IsAutoRenew int32 `json:"is_auto_renew"`

	// 实例名称
	Name string `json:"name"`

	// 网卡信息
	Nics []Nic `json:"nics"`

	OutsideInsConfig *OutsideInsConfig `json:"outside_ins_config,omitempty"`

	// 订购周期数目
	PeriodNum int32 `json:"period_num"`

	// 订购周期类型
	PeriodType int32 `json:"period_type"`

	// 产品列表
	ProductInfos []ProductInfoBeanNew `json:"product_infos"`

	// 促销ID
	PromotionActivityId *string `json:"promotion_activity_id,omitempty"`

	// 折扣信息
	PromotionInfo *string `json:"promotion_info,omitempty"`

	PublicIp *PublicIp `json:"public_ip,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 安全组信息
	SecurityGroups []SecurityGroup `json:"security_groups"`

	// 订购数量，当前仅支持1台
	SubscriptionNum int32 `json:"subscription_num"`

	// VPC ID
	VpcId string `json:"vpc_id"`
}

func (o ChargeOrderDbssNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeOrderDbssNew struct{}"
	}

	return strings.Join([]string{"ChargeOrderDbssNew", string(data)}, " ")
}
