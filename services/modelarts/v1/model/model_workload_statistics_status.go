package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadStatisticsStatus **参数描述**：不同状态下作业个数。
type WorkloadStatisticsStatus struct {

	// **参数描述**： 排队中的作业个数。 **取值范围**： 不涉及。
	Queue *int32 `json:"Queue,omitempty"`

	// **参数描述**： 等待中的作业个数。 **取值范围**： 不涉及。
	Pending *int32 `json:"Pending,omitempty"`

	// **参数描述**： 异常的作业个数。 **取值范围**： 不涉及。
	Abnormal *int32 `json:"Abnormal,omitempty"`

	// **参数描述**： 终止中的作业个数。 **取值范围**： 不涉及。
	Terminating *int32 `json:"Terminating,omitempty"`

	// **参数描述**： 创建中的作业个数。 **取值范围**： 不涉及。
	Creating *int32 `json:"Creating,omitempty"`

	// **参数描述**： 运行中的作业个数。 **取值范围**： 不涉及。
	Running *int32 `json:"Running,omitempty"`

	// **参数描述**： 已完成的作业个数。 **取值范围**： 不涉及。
	Completed *int32 `json:"Completed,omitempty"`

	// **参数描述**： 已终止的作业个数。 **取值范围**： 不涉及。
	Terminated *int32 `json:"Terminated,omitempty"`

	// **参数描述**：运行失败的作业个数。 **取值范围**： 不涉及。
	Failed *int32 `json:"Failed,omitempty"`
}

func (o WorkloadStatisticsStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadStatisticsStatus struct{}"
	}

	return strings.Join([]string{"WorkloadStatisticsStatus", string(data)}, " ")
}
