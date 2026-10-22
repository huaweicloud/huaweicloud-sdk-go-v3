package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComputeAllocatable **参数解释**：可分配资源汇总。
type ComputeAllocatable struct {

	// **参数解释**：cpu剩余量，单位毫核。 **取值范围**：小于等于资源池该规格cpu大小。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：内存剩余量，单位MB。 **取值范围**：小于等于资源池该规格内存大小。
	Mem *int32 `json:"mem,omitempty"`

	// **参数解释**：显卡剩余量，单位卡。 **取值范围**：小于等于资源池该规格显卡大小。
	Accelerator *int32 `json:"accelerator,omitempty"`
}

func (o ComputeAllocatable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeAllocatable struct{}"
	}

	return strings.Join([]string{"ComputeAllocatable", string(data)}, " ")
}
