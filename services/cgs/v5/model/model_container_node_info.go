package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点（服务器）状态列表
type ContainerNodeInfo struct {
	// 节点id

	HostId *string `json:"host_id,omitempty"`
	// 节点（服务器）名称

	HostName *string `json:"host_name,omitempty"`
	// 服务器状态，包含如下4种。   - ACTIVE ：正在运行。   - SHUTOFF ：关机。   - BUILDING ：创建中。   - ERROR ：故障。

	HostStatus *string `json:"host_status,omitempty"`
	// Agent状态，包含如下3种。   - not_register ：未注册。   - online ：在线。   - offline ：离线。

	AgentStatus *string `json:"agent_status,omitempty"`
	// 防护状态，包含如下2种。   - closed ：关闭。   - opened ：开启。

	ProtectStatus *string `json:"protect_status,omitempty"`
}

func (o ContainerNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerNodeInfo struct{}"
	}

	return strings.Join([]string{"ContainerNodeInfo", string(data)}, " ")
}
