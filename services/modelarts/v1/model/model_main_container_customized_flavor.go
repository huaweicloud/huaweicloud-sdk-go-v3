package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MainContainerCustomizedFlavor 配置训练作业自定义规格。
type MainContainerCustomizedFlavor struct {

	// **参数解释**：cpu核数。 **取值范围**：大于零。
	CpuCoreNum *float32 `json:"cpu_core_num,omitempty"`

	// **参数解释**：内存大小。 **取值范围**：大于零。
	MemSize *float32 `json:"mem_size,omitempty"`

	// **参数解释**：加速卡卡数。 **取值范围**：大于等于零。
	AcceleratorNum *float32 `json:"accelerator_num,omitempty"`
}

func (o MainContainerCustomizedFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MainContainerCustomizedFlavor struct{}"
	}

	return strings.Join([]string{"MainContainerCustomizedFlavor", string(data)}, " ")
}
