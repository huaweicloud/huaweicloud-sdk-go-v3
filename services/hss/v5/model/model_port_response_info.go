package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PortResponseInfo 端口信息
type PortResponseInfo struct {

	// **参数解释**: 主机ID **取值范围**: 字符长度1-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 监听ip **取值范围**: 字符长度1-128位
	Laddr *string `json:"laddr,omitempty"`

	// **参数解释**: 端口状态 **取值范围**: normal: 正常 danger: 危险 unknown: 未知
	Status *string `json:"status,omitempty"`

	// **参数解释**: 端口号 **取值范围**: 最小值0，最大值65535
	Port *int32 `json:"port,omitempty"`

	// **参数解释**: 端口类型：目前包括TCP，UDP两种 **取值范围**: TCP: TCP类型的端口 UDP: UDP类型的端口
	Type *string `json:"type,omitempty"`

	// **参数解释**: 进程ID **取值范围**: 最小值1，最大值65535
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**: 进程可执行文件路径 **取值范围**: 字符长度1-256位
	Path *string `json:"path,omitempty"`

	// **参数解释**: Agent ID **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**: 容器 ID **取值范围**: 字符长度0-128位
	ContainerId *string `json:"container_id,omitempty"`
}

func (o PortResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortResponseInfo struct{}"
	}

	return strings.Join([]string{"PortResponseInfo", string(data)}, " ")
}
