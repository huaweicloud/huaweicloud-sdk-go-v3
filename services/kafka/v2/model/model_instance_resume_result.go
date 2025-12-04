package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceResumeResult struct {

	// **参数解释**： 实例ID。   **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 任务ID。  **取值范围**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 错误信息。 **取值范围**： 不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o InstanceResumeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceResumeResult struct{}"
	}

	return strings.Join([]string{"InstanceResumeResult", string(data)}, " ")
}
