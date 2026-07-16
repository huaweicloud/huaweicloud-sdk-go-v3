package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeVoAnnotations 节点annotations
type NodeVoAnnotations struct {

	// **参数解释**：NPU卡的资源使用拓扑信息，长度为16的二进制编码，右起第一位编码代表卡1。其中，1表示占用，0表示空闲。例如，16卡的机型中卡1和卡15被占用，值为0100000000000001；8卡的机型中卡1和卡7被占用，返回值为0000000001000001。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsNpuTopologyPlacement *string `json:"os.modelarts/npu-topology-placement,omitempty"`
}

func (o NodeVoAnnotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeVoAnnotations struct{}"
	}

	return strings.Join([]string{"NodeVoAnnotations", string(data)}, " ")
}
