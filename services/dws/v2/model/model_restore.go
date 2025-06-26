package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Restore **参数解释**： 恢复对象。 **取值范围**： 不涉及。
type Restore struct {

	// **参数解释**： 集群名称。 **取值范围**： 要求唯一性，必须以字母开头并只包含字母、数字、中划线，下划线，长度为4~64个字符。
	Name string `json:"name"`

	// **参数解释**： 指定子网ID，用于集群网络配置。 **取值范围**： 默认值与原集群相同。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**： 指定安全组ID，用于集群网络配置。默认值与原集群相同。 **取值范围**： 不涉及。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数解释**： 指定虚拟私有云ID，用于集群网络配置。默认值与原集群相同。 **取值范围**： 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： 指定集群可用区。默认值与原集群相同。 **取值范围**： 不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**： 指定集群服务端口。 **取值范围**： 不涉及。
	Port *int32 `json:"port,omitempty"`

	PublicIp *PublicIp `json:"public_ip,omitempty"`

	// **参数解释**： 企业项目ID，对集群指定企业项目。如果未指定，则使用默认企业项目“default”的ID，即0。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 指定网络协议类型，表明是否支持IPv6，默认不使用IPv6。 **取值范围**： 不涉及。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`
}

func (o Restore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Restore struct{}"
	}

	return strings.Join([]string{"Restore", string(data)}, " ")
}
