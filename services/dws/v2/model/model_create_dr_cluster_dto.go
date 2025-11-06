package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrClusterDto **参数解释**： 备集群obs桶。 **取值范围**： 不涉及。
type CreateDrClusterDto struct {

	// **参数解释**： 集群名称。 **取值范围**： 不涉及。
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 子网的网段。 **取值范围**： 不涉及。
	Cidr *string `json:"cidr,omitempty"`

	// **参数解释**： 子网的网关。 **取值范围**： 不涉及。
	GatewayIp *string `json:"gateway_ip,omitempty"`

	// **参数解释**： 主网卡IP。 **取值范围**： 不涉及。
	InternalIps *string `json:"internal_ips,omitempty"`

	// **参数解释**： 内网IP。 **取值范围**： 不涉及。
	PrivateIps *string `json:"private_ips,omitempty"`

	// **参数解释**： 域名。 **取值范围**： 不涉及。
	Endpoint *string `json:"endpoint,omitempty"`

	// **参数解释**： 数据库管理员密码。 **取值范围**： 不涉及。
	DbAdminPwd *string `json:"db_admin_pwd,omitempty"`

	// **参数解释**： 容灾ID。 **取值范围**： 不涉及。
	DisasterRecoveryId *string `json:"disaster_recovery_id,omitempty"`

	// **参数解释**： 内核的版本。 **取值范围**： 不涉及。
	KernelVersion *string `json:"kernel_version,omitempty"`
}

func (o CreateDrClusterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrClusterDto struct{}"
	}

	return strings.Join([]string{"CreateDrClusterDto", string(data)}, " ")
}
