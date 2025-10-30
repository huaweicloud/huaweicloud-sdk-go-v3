package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelModuleHostInfo **参数解释** 内核模块服务器列表
type KernelModuleHostInfo struct {

	// **参数解释** agent id **取值范围** 字符长度1-64
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释** 主机id **取值范围** 字符长度1-64
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 服务器名称 **取值范围** 字符长度1-64
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 服务器ip **取值范围** 字符长度1-64
	HostIp *string `json:"host_ip,omitempty"`

	KernelModuleInfo *KernelModuleInfo `json:"kernel_module_info,omitempty"`
}

func (o KernelModuleHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KernelModuleHostInfo struct{}"
	}

	return strings.Join([]string{"KernelModuleHostInfo", string(data)}, " ")
}
