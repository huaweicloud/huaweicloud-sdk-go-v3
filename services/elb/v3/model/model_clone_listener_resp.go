package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneListenerResp clone出的监听器信息。
type CloneListenerResp struct {

	// **参数解释**：新监听器ID。 **取值范围**：标准的UUID格式，长度为36个字符。
	Id *string `json:"id,omitempty"`

	// **参数解释**：目的负载均衡器ID。 **取值范围**：标准的UUID格式，长度为36个字符。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// **参数解释**：新监听器监听端口。 **取值范围**：1-65535
	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	// **参数解释**：端口监听范围（闭区间)。
	PortRanges *[]ResPortRange `json:"port_ranges,omitempty"`
}

func (o CloneListenerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneListenerResp struct{}"
	}

	return strings.Join([]string{"CloneListenerResp", string(data)}, " ")
}
