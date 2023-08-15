package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigCommodityOrder 获取ApigCommodityOrder实例列表信息
type ApigCommodityOrder struct {

	// 实例所属项目id
	ProjectId *string `json:"project_id,omitempty"`

	// CBC订单id
	OrderId *string `json:"order_id,omitempty"`

	// 当前所属region Id
	RegionId *string `json:"region_id,omitempty"`

	// 实例id
	ResourceId *string `json:"resource_id,omitempty"`

	// 实例名称
	ResourceName *string `json:"resource_name,omitempty"`

	// CBC订购id
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// 资源类型，hws.resource.type.dayu
	ResourceType *string `json:"resource_type,omitempty"`

	// 产品规格编码，例如dayu.starter，dayu.basic，dayu.advanced等
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// CBC产品id，未安装CBC的环境无需这个值
	ProductId *string `json:"product_id,omitempty"`

	// 订单类型标识符
	OrderType *string `json:"order_type,omitempty"`

	// 支付选项，留空
	ChargeType *string `json:"charge_type,omitempty"`

	// 自动续费标识，当前实例为按需支付时必填，0代表不续费，1代表自动续费
	IsAutoRenew *int32 `json:"is_auto_renew,omitempty"`

	// 实例状态,1未生效2生效中3已删除=退订4保留期=冻结5宽限期6删除中
	Status *int32 `json:"status,omitempty"`

	// 虚拟私有云id
	VpcId *string `json:"vpc_id,omitempty"`

	// 安全组id
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 企业项目id，如果当前为公有云，且用户开启企业项目，则必选
	EpsId *string `json:"eps_id,omitempty"`

	// 生效时间点，包周期实例有效
	EffectiveTime float32 `json:"effective_time,omitempty"`

	// 过期时间天数，包周期实例有效
	ExpireDays *string `json:"expire_days,omitempty"`

	// 过期时间点，包周期有效
	ExpireTime float32 `json:"expire_time,omitempty"`

	// CBC锁定节点
	LockCheckEndpoint *string `json:"lock_check_endpoint,omitempty"`

	// 创建用户
	CreateUser *string `json:"create_user,omitempty"`

	// 创建时间点
	CreateTime float32 `json:"create_time,omitempty"`

	// 用户domain id
	DomainId *string `json:"domain_id,omitempty"`

	// 是否试用订单
	IsTrialOrder *int32 `json:"is_trial_order,omitempty"`

	// 工作空间模式说明
	WorkSpaceMode *string `json:"work_space_mode,omitempty"`
}

func (o ApigCommodityOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigCommodityOrder struct{}"
	}

	return strings.Join([]string{"ApigCommodityOrder", string(data)}, " ")
}
