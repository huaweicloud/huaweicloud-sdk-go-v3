package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListenerNode
type ListenerNode struct {

	// **参数解释**：监听器id。  **取值范围**：不涉及
	Id string `json:"id"`

	// **参数解释**：监听器名称。  **取值范围**：不涉及
	Name string `json:"name"`

	// **参数解释**：监听器协议。  **取值范围**：不涉及
	Protocol string `json:"protocol"`

	// **参数解释**：监听器监听端口。  **取值范围**：不涉及
	ProtocolPort int32 `json:"protocol_port"`

	// **参数解释**：全端口监听，指定端口监听范围（闭区间)，最多指定10个端口组，每个组范围不可有重叠部分 >仅当protocol_port为0时可以传入。
	PortRanges *[]PortRange `json:"port_ranges,omitempty"`
}

func (o ListenerNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerNode struct{}"
	}

	return strings.Join([]string{"ListenerNode", string(data)}, " ")
}
