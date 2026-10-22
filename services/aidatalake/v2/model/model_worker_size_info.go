package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkerSizeInfo **参数解释**：可选资源配置枚举。
type WorkerSizeInfo struct {

	// **参数解释**：可选资源配置项中的cpu大小，即单个worker的cpu配置大小，单位毫核。 **取值范围**：小于等于资源池该规格cpu大小。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：可选资源配置项中的内存大小，即单个worker的内存配置大小，单位MB。 **取值范围**：小于等于资源池该规格内存大小。
	Mem *int32 `json:"mem,omitempty"`

	// **参数解释**：可选资源配置项中的显卡数量，即单个worker的显卡配置数量，单位卡。 **取值范围**：小于等于资源池该规格显卡大小。
	Accelerator *int32 `json:"accelerator,omitempty"`
}

func (o WorkerSizeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkerSizeInfo struct{}"
	}

	return strings.Join([]string{"WorkerSizeInfo", string(data)}, " ")
}
