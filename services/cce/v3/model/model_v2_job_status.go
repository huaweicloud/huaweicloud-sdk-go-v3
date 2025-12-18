package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type V2JobStatus struct {

	// **参数解释**： Job的状态 **约束限制**： 不涉及 **取值范围**： - Initializing：未执行 - Running：执行中 - Failed：执行失败 - Success：执行成功  **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： Job执行失败的原因 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Reason *string `json:"reason,omitempty"`

	// **参数解释**： Job完成时间 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CompletionTime *string `json:"completionTime,omitempty"`
}

func (o V2JobStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2JobStatus struct{}"
	}

	return strings.Join([]string{"V2JobStatus", string(data)}, " ")
}
