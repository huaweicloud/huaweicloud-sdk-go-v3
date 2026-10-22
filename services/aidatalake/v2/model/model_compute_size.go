package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComputeSize **参数解释**：资源总量信息。
type ComputeSize struct {

	// **参数解释**：cpu总量，单位毫核。 **取值范围**：等于资源池该规格cpu大小。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：内存总量，单位MB。 **取值范围**：等于资源池该规格内存大小。
	Mem *int32 `json:"mem,omitempty"`

	// **参数解释**：显卡总量，单位卡。 **取值范围**：等于资源池该规格显卡数量大小。
	Accelerator *int32 `json:"accelerator,omitempty"`
}

func (o ComputeSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeSize struct{}"
	}

	return strings.Join([]string{"ComputeSize", string(data)}, " ")
}
