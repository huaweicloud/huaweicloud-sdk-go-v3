package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkloadNodeVo struct {

	// **参数解释**：作业运行节点的IP地址。 **取值范围**：不涉及。
	HostIp *string `json:"hostIp,omitempty"`

	// **参数解释**：NPU卡的资源使用拓扑信息，长度为16的二进制编码，右起第一位编码代表卡1。其中，1表示占用，0表示空闲。例如，16卡的机型中卡1和卡15被占用，值为0100000000000001；8卡的机型中卡1和卡7被占用，返回值为0000000001000001。 **取值范围**：不涉及。
	NpuTopologyPlacement *string `json:"npuTopologyPlacement,omitempty"`

	ResourceRequirement *ResourceRequirementVo `json:"resourceRequirement,omitempty"`
}

func (o WorkloadNodeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadNodeVo struct{}"
	}

	return strings.Join([]string{"WorkloadNodeVo", string(data)}, " ")
}
