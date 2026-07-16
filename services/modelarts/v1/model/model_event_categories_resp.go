package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventCategoriesResp 事件类型
type EventCategoriesResp struct {

	// **参数解释**：事件类型id。 **取值范围**：枚举值如下： - JobStarted：作业开始 - JobCompleted：作业结束 - JobFailed：作业失败 - JobTerminated：作业终止 - JobRestarted：作业重启 - JobHanged：作业疑似卡死 - JobPreempted：作业抢占
	Id *string `json:"id,omitempty"`

	// **参数解释**：事件类型名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：事件类型描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：是否通知。 **取值范围**： - true：通知 - false：不通知
	Notification *bool `json:"notification,omitempty"`
}

func (o EventCategoriesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventCategoriesResp struct{}"
	}

	return strings.Join([]string{"EventCategoriesResp", string(data)}, " ")
}
