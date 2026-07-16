package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowedAddressPairs **参数解释：** IP/Mac地址对。 **约束限制：** - IP地址不允许为 “0.0.0.0/0”。 - 如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。 - 如果allowed_address_pairs为“1.1.1.1/0”，表示关闭源目的地址检查开关。 - 如果是虚拟IP绑定云服务器    则mac_address可为空或者填写被绑定云服务器网卡的Mac地址。    被绑定的云服务器网卡allowed_address_pairs的IP地址填“1.1.1.1/0”。 **取值范围：** 不涉及 **默认取值：** 不涉及
type AllowedAddressPairs struct {

	// **参数解释：** IP地址。 **约束限制：** - 不支持0.0.0.0/0。 - 如果allowed_address_pairs配置地址池较大的CIDR（掩码小于24位），建议为该port配置一个单独的安全组。 **取值范围：** 不涉及 **默认取值：** 不涉及
	IpAddress *string `json:"ip_address,omitempty"`

	// **参数解释：** MAC地址。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	MacAddress *string `json:"mac_address,omitempty"`
}

func (o AllowedAddressPairs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowedAddressPairs struct{}"
	}

	return strings.Join([]string{"AllowedAddressPairs", string(data)}, " ")
}
