package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterSecurityCheckPortInfo 端口信息
type ClusterSecurityCheckPortInfo struct {

	// **参数解释**： 端口号 **取值范围**： 不涉及
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 容器ID **取值范围**： 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器名称 **取值范围**： 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**： pod名称 **取值范围**： 不涉及
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**： 节点名称 **取值范围**： 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 私有IP地址 **取值范围**： 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 公有IP地址 **取值范围**： 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 程序文件 **取值范围**： 不涉及
	Path *string `json:"path,omitempty"`

	// **参数解释**： 进程PID **取值范围**： 不涉及
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**： 监听IP **取值范围**： 不涉及
	Laddr *string `json:"laddr,omitempty"`
}

func (o ClusterSecurityCheckPortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSecurityCheckPortInfo struct{}"
	}

	return strings.Join([]string{"ClusterSecurityCheckPortInfo", string(data)}, " ")
}
