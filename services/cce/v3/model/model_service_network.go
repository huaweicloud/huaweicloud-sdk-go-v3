package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceNetwork struct {

	// **参数解释：** kubernetes clusterIP IPv4 CIDR取值范围。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 默认为\"10.247.0.0/16\"。
	IPv4CIDR *string `json:"IPv4CIDR,omitempty"`

	// **参数解释：** kubernetes clusterIP IPv6 CIDR取值范围。 **约束限制：** 仅开启IPV6双栈的Turbo集群支持配置IPv6服务网段。 **取值范围：** 不涉及 **默认取值：** Turbo集群默认为\"fc00::/112\" CCE集群默认为\"fd00:1234::/120\"
	IPv6CIDR *string `json:"IPv6CIDR,omitempty"`
}

func (o ServiceNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceNetwork struct{}"
	}

	return strings.Join([]string{"ServiceNetwork", string(data)}, " ")
}
