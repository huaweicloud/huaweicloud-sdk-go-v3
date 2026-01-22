package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceItem struct {

	// **参数解释**： 协议类型，用于明确规则网络协议 **约束限制**：  RuleServiceDto.type为0时，此处不能为空。 **取值范围**： - 6：TCP - 17：UDP - 1：ICMP - 58：ICMPV6 - -1：Any **默认取值**： 不涉及
	Protocol *int32 `json:"protocol,omitempty"`

	// **参数解释**： 源端口，会话发起方的端口。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SourcePort *string `json:"source_port,omitempty"`

	// **参数解释**： 目的端口，会话接收方的端口。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	DestPort *string `json:"dest_port,omitempty"`

	// **参数解释**： 服务组成员描述信息 **约束限制**： 字符串长度0-255 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 服务（协议、源端口、目的端口）成员 **约束限制**： 字符串长度0-255 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o ServiceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceItem struct{}"
	}

	return strings.Join([]string{"ServiceItem", string(data)}, " ")
}
