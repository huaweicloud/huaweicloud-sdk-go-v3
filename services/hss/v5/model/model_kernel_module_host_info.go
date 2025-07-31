package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KernelModuleHostInfo 内核模块服务器列表
type KernelModuleHostInfo struct {

	// agent_id
	AgentId *string `json:"agent_id,omitempty"`

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ip
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
