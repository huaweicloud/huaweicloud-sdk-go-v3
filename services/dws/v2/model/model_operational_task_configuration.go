package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationalTaskConfiguration struct {

	// **参数解释**： 用户表VacuumFull运维任务最小并发数。 **约束限制**： 不涉及。 **取值范围**： 1~24 **默认取值**： 不涉及。
	ParallelMin *int32 `json:"parallel_min,omitempty"`

	// **参数解释**： 用户表VacuumFull运维任务最大并发数。 **约束限制**： 不涉及。 **取值范围**： 1~24 **默认取值**： 不涉及。
	ParallelMax *int32 `json:"parallel_max,omitempty"`

	// **参数解释**： 小CU阈值。 **约束限制**： 不涉及。 **取值范围**： 0~1000。 **默认取值**： 不涉及。
	SmallCuRowsLimit *int32 `json:"small_cu_rows_limit,omitempty"`

	// **参数解释**： 小CU占比。 **约束限制**： 不涉及。 **取值范围**： 0.01~99.99。 **默认取值**： 不涉及。
	SmallCuPercentage *float64 `json:"small_cu_percentage,omitempty"`
}

func (o OperationalTaskConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationalTaskConfiguration struct{}"
	}

	return strings.Join([]string{"OperationalTaskConfiguration", string(data)}, " ")
}
