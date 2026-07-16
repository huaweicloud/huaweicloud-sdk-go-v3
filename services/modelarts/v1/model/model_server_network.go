package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerNetwork struct {

	// **参数解释**：创建服务器是否启用IPv6。表示在创建服务器时是否启用IPv6支持。 **约束限制**：不涉及。 **取值范围**： - true：启用IPv6 - false：不启用IPv6 **默认取值**：不涉及。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// **参数解释**：服务器RoCE网络ID。表示服务器的RoCE网络ID。 **约束限制**：不涉及。 **取值范围**：必须是UUID格式的字符串。 **默认取值**：不涉及。
	RoceId *string `json:"roce_id,omitempty"`

	// **参数解释**：服务器所在的安全组ID。表示服务器所属的安全组ID。 **约束限制**：不涉及。 **取值范围**：必须是UUID格式的字符串。 **默认取值**：不涉及。
	SecurityGroupId string `json:"security_group_id"`

	// **参数解释**：服务器所在子网ID。表示服务器所属的子网ID。 **约束限制**：不涉及。 **取值范围**：必须是UUID格式的字符串。 **默认取值**：不涉及。
	SubnetId string `json:"subnet_id"`

	// **参数解释**：服务器所在虚拟私有云ID。表示服务器所属的虚拟私有云ID。 **约束限制**：不涉及。 **取值范围**：必须是UUID格式的字符串。 **默认取值**：不涉及。
	VpcId string `json:"vpc_id"`

	HyperCluster *ServerNetworkHyperCluster `json:"hyper_cluster,omitempty"`

	// **参数解释：** IP/Mac对列表。 **约束限制：** - IP地址不允许为 “0.0.0.0/0”。 - 如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。 - 如果allowed_address_pairs为“1.1.1.1/0”，表示关闭源目的地址检查开关。 - 如果是虚拟IP绑定云服务器，       则mac_address可为空或者填写被绑定云服务器网卡的Mac地址。       被绑定的云服务器网卡allowed_address_pairs的IP地址填“1.1.1.1/0”。 **取值范围：** 不涉及 **默认取值：** 不涉及\"
	AllowedAddressPairs *[]AllowedAddressPairs `json:"allowed_address_pairs,omitempty"`
}

func (o ServerNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerNetwork struct{}"
	}

	return strings.Join([]string{"ServerNetwork", string(data)}, " ")
}
