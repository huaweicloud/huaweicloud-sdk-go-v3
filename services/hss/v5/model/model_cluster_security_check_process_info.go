package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterSecurityCheckProcessInfo 进程信息
type ClusterSecurityCheckProcessInfo struct {

	// **参数解释**： 容器ID **取值范围**： 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器名称 **取值范围**： 不涉及
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**： 节点名称 **取值范围**： 不涉及
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 私有IP地址 **取值范围**： 不涉及
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 公有IP地址 **取值范围**： 不涉及
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 启动时间 **取值范围**： 不涉及
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 进程PID **取值范围**： 不涉及
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**： 文件权限 **取值范围**： 不涉及
	Permission *string `json:"permission,omitempty"`

	// **参数解释**： 运行用户 **取值范围**： 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 启动参数 **取值范围**： 不涉及
	LaunchParams *string `json:"launch_params,omitempty"`

	// **参数解释**： 文件hash **取值范围**： 不涉及
	Hash *string `json:"hash,omitempty"`
}

func (o ClusterSecurityCheckProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSecurityCheckProcessInfo struct{}"
	}

	return strings.Join([]string{"ClusterSecurityCheckProcessInfo", string(data)}, " ")
}
