package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfoContainerForensic **参数解释**： 容器类型告警取证信息 **取值范围**： 不涉及
type EventForensicInfoContainerForensic struct {

	// **参数解释**： 容器ID **取值范围**： 不涉及
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器状态 **取值范围**： 不涉及
	ContainerStatus *string `json:"container_status,omitempty"`

	// **参数解释**： pod uid **取值范围**： 不涉及
	PodUid *string `json:"pod_uid,omitempty"`

	// **参数解释**： pod name **取值范围**： 不涉及
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释**： 命名空间 **取值范围**： 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 集群id **取值范围**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 镜像ID **取值范围**： 不涉及
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**： 镜像名称 **取值范围**： 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**： 容器特权开放 **取值范围**： 不涉及
	Priviledged *bool `json:"priviledged,omitempty"`

	// **参数解释**： 容器pid命名空间模式 **取值范围**： 不涉及
	PidMode *string `json:"pid_mode,omitempty"`

	// **参数解释**： 容器ipc命名空间模式 **取值范围**： 不涉及
	IpcMode *string `json:"ipc_mode,omitempty"`

	// **参数解释**： 容器网络命名空间模式 **取值范围**： 不涉及
	NetworkMode *string `json:"network_mode,omitempty"`

	// **参数解释**： 容器开放所有端口 **取值范围**： 不涉及
	PublishAllports *bool `json:"publish_allports,omitempty"`

	// **参数解释**： 容器权限 **取值范围**： 不涉及
	Capabilities *string `json:"capabilities,omitempty"`

	// **参数解释**： 容器安全选项 **取值范围**： 不涉及
	SecurityOpts *string `json:"security_opts,omitempty"`

	// **参数解释**： 容器开放端口 **取值范围**： 不涉及
	Ports *string `json:"ports,omitempty"`

	// **参数解释**： 容器挂载目录 **取值范围**： 不涉及
	Mounts *string `json:"mounts,omitempty"`

	// **参数解释**： 系统调用 **取值范围**： 不涉及
	SysCall *string `json:"sys_call,omitempty"`

	// **参数解释**： 容器名称 **取值范围**： 不涉及
	ContainerName *string `json:"container_name,omitempty"`
}

func (o EventForensicInfoContainerForensic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfoContainerForensic struct{}"
	}

	return strings.Join([]string{"EventForensicInfoContainerForensic", string(data)}, " ")
}
