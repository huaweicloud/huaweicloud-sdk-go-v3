package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSpec **参数解释**：资源配置规格，用于指定作业使用CPU、内存、磁盘资源，可生成对Driver、Executor等资源的粗粒度描述。 **约束限制**：不涉及。
type ResourceSpec struct {

	// **参数解释**：CPU核数，单位为毫核，用于指定作业运行所需的CPU资源。 **约束限制**：不涉及。 **取值范围**：最小值为1000，最大值为48000。 **默认取值**：1000。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：内存大小，单位为MB，用于指定作业运行所需的内存资源。 **约束限制**：不涉及。 **取值范围**：最小值为1024，最大值为393216。 **默认取值**：4096。
	Memory *int32 `json:"memory,omitempty"`

	// **参数解释**：本地磁盘大小，用于指定Spark作业Driver和Executor的本地磁盘资源。 **约束限制**：不涉及。 **取值范围**：单位为GB，最小值为0，最大值为994。 **默认取值**：不涉及。
	Disk *int32 `json:"disk,omitempty"`
}

func (o ResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpec struct{}"
	}

	return strings.Join([]string{"ResourceSpec", string(data)}, " ")
}
