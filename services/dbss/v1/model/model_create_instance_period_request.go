package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstancePeriodRequest struct {

	// 实例名称。取值范围： - 只能由中文字符、英文字母、数字、下划线、中划线组成，且长度小于等于64个字符。
	Name string `json:"name"`

	// 云服务器使用的规格ID
	FlavorRef string `json:"flavor_ref"`

	// 虚拟私有云的ID
	VpcId string `json:"vpc_id"`

	// 云服务器对应可用分区信息。(两个主备分区，中间用“,”分割，例如az1.dc1,az2.dc2)。
	AvailabilityZone string `json:"availability_zone"`

	// 企业项目ID。对接EPS必输。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 云服务器对应的网卡信息
	Nics []CreateInstancePeriodRequestNics `json:"nics"`

	// 云服务器对应安全组信息
	SecurityGroups []CreateInstancePeriodRequestSecurityGroups `json:"security_groups"`

	// 备注信息
	Comment *string `json:"comment,omitempty"`

	// 云服务器所在区域ID
	Region string `json:"region"`

	// 服务类型： - hws.service.type.dbss
	CloudServiceType string `json:"cloud_service_type"`

	// 计费模式： - 0: 包周期计费 - 1: 按需计费
	ChargingMode int32 `json:"charging_mode"`

	// -订购周期类型 - 0: 天 - 1：周 - 2：月 - 3：年 - 4: 小时 - 5: 绝对时间
	PeriodType int32 `json:"period_type"`

	// 订购周期数
	PeriodNum int32 `json:"period_num"`

	// 订购数量： DBSS只支持订购1套，不支持多套
	SubscriptionNum int32 `json:"subscription_num"`

	// 产品信息列表
	ProductInfos []CreateInstancePeriodRequestProductInfos `json:"product_infos"`

	// 资源标签
	Tags *[]KeyValueBean `json:"tags,omitempty"`

	// 折扣信息
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 自动续费 - 1: 自动续费 - 0: 不自动续费
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`
}

func (o CreateInstancePeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancePeriodRequest struct{}"
	}

	return strings.Join([]string{"CreateInstancePeriodRequest", string(data)}, " ")
}
