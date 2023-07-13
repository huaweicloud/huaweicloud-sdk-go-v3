package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterElbInfo Elb返回体
type ClusterElbInfo struct {

	// 弹性负载均衡ID
	Id *string `json:"id,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 弹性负载均衡名称
	Name *string `json:"name,omitempty"`

	// 弹性负载均衡描述
	Description *string `json:"description,omitempty"`

	// 弹性负载均衡地址
	VipAddress *string `json:"vip_address,omitempty"`

	// 子网ID
	VipSubnetId *string `json:"vip_subnet_id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 弹性负载均衡类型。枚举值：Internal,External
	Type *string `json:"type,omitempty"`

	// 弹性负载均衡的管理状态。枚举值：ACTIVE,PENDING_CREATE,ERROR
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 绑定状态： 0为未绑定，1为已绑定
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 虚拟私有云ID
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o ClusterElbInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterElbInfo struct{}"
	}

	return strings.Join([]string{"ClusterElbInfo", string(data)}, " ")
}
