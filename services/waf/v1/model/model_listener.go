package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Listener **参数解释：** 监听器列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type Listener struct {

	// **参数解释：** 监听器的名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 监听器的ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 监听器的监听协议 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Protocol *string `json:"protocol,omitempty"`

	// **参数解释：** 监听器的监听端口 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ProtocolPort *int32 `json:"protocol_port,omitempty"`
}

func (o Listener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Listener struct{}"
	}

	return strings.Join([]string{"Listener", string(data)}, " ")
}
