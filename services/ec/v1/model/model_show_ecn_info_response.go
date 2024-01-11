package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEcnInfoResponse Response Object
type ShowEcnInfoResponse struct {

	// 企业连接网络ID
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 租户账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 企业连接网络名字
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 带宽
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 企业连接网络类型
	Type *string `json:"type,omitempty"`

	// 大区ID
	AreaId *string `json:"area_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 企业连接网络AS号
	EcnAsn *int64 `json:"ecn_asn,omitempty"`

	// 智能企业网关AS号
	IegAsn *int64 `json:"ieg_asn,omitempty"`

	// VXLAN网络标识
	Vni *int32 `json:"vni,omitempty"`

	// 企业路由器列表
	EnterpriseRouterIds *[]string `json:"enterprise_router_ids,omitempty"`

	// 虚拟私有云列表
	VpcIds *[]string `json:"vpc_ids,omitempty"`

	// 绑定智能企业网关数量
	BindIegCount *int32 `json:"bind_ieg_count,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 冻结效果
	FrozenEffect *int32 `json:"frozen_effect,omitempty"`

	// 分支互联开关
	HubEnable *bool `json:"hub_enable,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 包周期场景下购买的订单号，按需场景下为空
	OrderId *string `json:"order_id,omitempty"`

	// 包周期场景下购买的订单号，按需场景下为空
	ProductId      *string `json:"product_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEcnInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEcnInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowEcnInfoResponse", string(data)}, " ")
}
