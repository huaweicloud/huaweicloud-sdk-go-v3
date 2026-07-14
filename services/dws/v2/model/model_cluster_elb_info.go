package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterElbInfo **参数解释**： 弹性负载均衡详情。 **取值范围**： 不涉及。
type ClusterElbInfo struct {

	// **参数解释**： 弹性负载均衡ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 弹性负载均衡名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 弹性负载均衡描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 弹性负载均衡地址。 **取值范围**： 不涉及。
	VipAddress *string `json:"vip_address,omitempty"`

	// **参数解释**： 子网ID。 **取值范围**： 不涉及。
	VipSubnetId *string `json:"vip_subnet_id,omitempty"`

	// **参数解释**： 租户ID。 **取值范围**： 不涉及。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释**： 弹性负载均衡类型。 **取值范围**： Internal：独享型。 External：共享型。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 弹性负载均衡的管理状态。 **取值范围**： true：使用中 false：未使用
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**： 带宽信息。 **取值范围**： 大于等于0的非负整数。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// **参数解释**： 虚拟私有云ID。 **取值范围**： 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： 是否支持跨vpc绑定。 **取值范围**： 不涉及。
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`
}

func (o ClusterElbInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterElbInfo struct{}"
	}

	return strings.Join([]string{"ClusterElbInfo", string(data)}, " ")
}
