package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Workload 作业详细信息。
type Workload struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v1：当前资源版本为v1
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - Workload：资源池作业
	Kind string `json:"kind"`

	// **参数解释**：资源池中作业的业务类型。 **取值范围**：可选值如下： - train：训练作业 - infer：推理服务 - notebook：Notebook作业 - x-infer：新版推理作业
	Type string `json:"type"`

	// **参数解释**：集群中作业所属的命名空间。 **取值范围**：不涉及。
	Namespace string `json:"namespace"`

	// **参数解释**：作业的名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：作业的归属的上层业务的名称。 **取值范围**：不涉及。
	JobName *string `json:"jobName,omitempty"`

	// **参数解释**：作业的ID。 **取值范围**：不涉及。
	Uid *string `json:"uid,omitempty"`

	// **参数解释**：作业的归属的上层业务的ID。 **取值范围**：不涉及。
	JobUUID *string `json:"jobUUID,omitempty"`

	// **参数解释**：作业的资源规格。 **取值范围**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：作业状态。 **取值范围**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：作业创建者。 **取值范围**：不涉及。
	UserName *string `json:"userName,omitempty"`

	ResourceRequirement *WorkloadResourceRequirement `json:"resourceRequirement,omitempty"`

	// **参数解释**：作业的优先级。 **取值范围**：不涉及。
	Priority *string `json:"priority,omitempty"`

	// **参数解释**：作业的运行时长，以秒为单位。 **取值范围**：不涉及。
	RunningDuration *int32 `json:"runningDuration,omitempty"`

	// **参数解释**：作业的排队时长，以秒为单位。 **取值范围**：不涉及。
	PendingDuration *int32 `json:"pendingDuration,omitempty"`

	// **参数解释**：作业当前的排队位置。 **取值范围**：不涉及。
	PendingPosition *int32 `json:"pendingPosition,omitempty"`

	// **参数解释**：作业的Unix创建时间戳，以毫秒为单位。 **取值范围**：不涉及。
	CreateTime *int32 `json:"createTime,omitempty"`

	// **参数解释**：作业的k8s资源类型、分组和版本。 **取值范围**：不涉及。
	Gvk *string `json:"gvk,omitempty"`

	// **参数解释**：作业运行的节点IP列表，以“,”分隔。 **取值范围**：不涉及。
	HostIps *string `json:"hostIps,omitempty"`

	// **参数解释**：作业运行时占用的节点资源信息。
	Nodes *[]WorkloadNodeVo `json:"nodes,omitempty"`
}

func (o Workload) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Workload struct{}"
	}

	return strings.Join([]string{"Workload", string(data)}, " ")
}
