package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkResourceSpec **参数解释**：资源规格，定义单个组件（JobManager或TaskManager）的CPU、内存和磁盘配额。 **约束限制**：不涉及。
type FlinkResourceSpec struct {

	// **参数解释**：CPU毫核数，单位为毫核，用于指定JobManager或TaskManager的CPU资源。 **约束限制**：不涉及。 **取值范围**：500~48000。 **默认取值**：1000。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：内存大小，单位为MB，用于指定JobManager或TaskManager的内存资源。 **约束限制**：不涉及。 **取值范围**：2048~393216。 **默认取值**：4096。
	Memory *int32 `json:"memory,omitempty"`

	// **参数解释**：Pod临时存储容量，单位为GB，用于存放作业运行时的临时文件，临时存储不持久化。 **约束限制**：不涉及。 **取值范围**：30~1024。 **默认取值**：30。
	Disk *int32 `json:"disk,omitempty"`
}

func (o FlinkResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkResourceSpec struct{}"
	}

	return strings.Join([]string{"FlinkResourceSpec", string(data)}, " ")
}
