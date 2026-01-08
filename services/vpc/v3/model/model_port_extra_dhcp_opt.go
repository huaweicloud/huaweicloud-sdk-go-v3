package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PortExtraDhcpOpt struct {

	// **参数解释**： DHCP属性名称。在DHCP服务器为客户端分配IP地址时，传递的额外的控制信息或网络配置参数的名称。 **取值范围**： 不涉及。
	OptName string `json:"opt_name"`

	// **参数解释**： DHCP属性值。在DHCP服务器为客户端分配IP地址时，传递的额外的控制信息或网络配置参数的值。 **取值范围**： 不涉及。
	OptValue string `json:"opt_value"`
}

func (o PortExtraDhcpOpt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortExtraDhcpOpt struct{}"
	}

	return strings.Join([]string{"PortExtraDhcpOpt", string(data)}, " ")
}
