package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EipInfo struct {

	// **参数解释**： 弹性公网ID，可通过调用弹性IP列表查询获取，通过返回值中的data.records.id（.表示各对象之间层级的区分）获得 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	EipId *string `json:"eip_id,omitempty"`

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// **参数解释**： 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志ID，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象ID，type为1的为VPC边界防护对象ID。 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	ObjectId *string `json:"object_id,omitempty"`

	// **参数解释**： IPV4地址  **约束限制**： 不涉及 **取值范围**： 不涉及   **默认取值**： 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： IPV6地址  **约束限制**： 不涉及 **取值范围**： 不涉及   **默认取值**： 不涉及
	PublicIpv6 *string `json:"public_ipv6,omitempty"`

	// **参数解释**： EIP白名单标志 **约束限制**： 不涉及 **取值范围**： 0表示是EIP白名单，1表示不是EIP白名单 **默认取值**： 不涉及
	Type *int32 `json:"type,omitempty"`
}

func (o EipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipInfo struct{}"
	}

	return strings.Join([]string{"EipInfo", string(data)}, " ")
}
