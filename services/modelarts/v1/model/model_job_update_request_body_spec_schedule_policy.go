package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobUpdateRequestBodySpecSchedulePolicy **参数解释**：训练作业调度策略。 **约束限制**：不涉及。
type JobUpdateRequestBodySpecSchedulePolicy struct {

	// **参数解释**：训练作业优先级。 **约束限制**： - 仅使用专属资源池训练时才支持设置训练作业优先级。 - 作业优先级取值为1~3，默认优先级为1，最高优先级为3。 默认用户权限可选择优先级1和2，配置了“设置作业为高优先级权限”的用户可选择优先级1~3。  **取值范围**：1-3 **默认取值**：不涉及。
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释**：是否可以被抢占。 **约束限制**：不涉及。 **取值范围**： - true：可以被抢占 - false：不可以被抢占  **默认取值**：不涉及。
	Preemptible *bool `json:"preemptible,omitempty"`
}

func (o JobUpdateRequestBodySpecSchedulePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobUpdateRequestBodySpecSchedulePolicy struct{}"
	}

	return strings.Join([]string{"JobUpdateRequestBodySpecSchedulePolicy", string(data)}, " ")
}
