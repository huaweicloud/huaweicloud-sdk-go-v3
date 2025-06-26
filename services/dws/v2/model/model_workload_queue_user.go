package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueUser **参数解释**： 工作队列用户。 **取值范围**： 不涉及。
type WorkloadQueueUser struct {

	// **参数解释**： 用户名。 **取值范围**： 不涉及。
	UserName string `json:"user_name"`

	// **参数解释**： 执行计划阶段。 **取值范围**： 不涉及。
	OccupyResourceList []OccupyResource `json:"occupy_resource_list"`

	// **参数解释**： 执行结果。 **取值范围**： 不涉及。
	ExecResult *int32 `json:"exec_result,omitempty"`

	// **参数解释**： 执行日志。 **取值范围**： 不涉及。
	ExecLog *string `json:"exec_log,omitempty"`
}

func (o WorkloadQueueUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueUser struct{}"
	}

	return strings.Join([]string{"WorkloadQueueUser", string(data)}, " ")
}
