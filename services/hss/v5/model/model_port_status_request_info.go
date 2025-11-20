package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PortStatusRequestInfo struct {

	// **参数解释**: agent id **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	AgentId string `json:"agent_id"`

	// **参数解释**: 本地端口号 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2147483647 **默认取值**: 不涉及
	Port int32 `json:"port"`

	// **参数解释**: 端口类型 **约束限制**: 不涉及 **取值范围**: - TCP - UDP  **默认取值**: 不涉及
	PortType string `json:"port_type"`

	// **参数解释**: 容器id **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 主机id **约束限制**: 不涉及 **取值范围**: 字符长度0-64位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`
}

func (o PortStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"PortStatusRequestInfo", string(data)}, " ")
}
