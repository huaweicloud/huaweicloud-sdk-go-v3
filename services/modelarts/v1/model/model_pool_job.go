package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolJob 资源管理作业。
type PoolJob struct {

	// **参数解释**： job开始处理时间，单位毫秒。 **取值范围**： 不涉及。
	StartTimestamp *int64 `json:"startTimestamp,omitempty"`

	// **参数解释**： Job结束时间，单位毫秒。 **取值范围**： 不涉及。
	EndTimestamp *int64 `json:"endTimestamp,omitempty"`

	// **参数解释**： 任务ID。 **取值范围**： 不涉及。
	JobId *string `json:"jobId,omitempty"`

	// **参数解释**： 任务名称。 **取值范围**： 不涉及。
	JobName *string `json:"jobName,omitempty"`

	// **参数解释**： Job关联的资源，比如资源池描述。 **取值范围**： 不涉及。
	InvolvedObjects *string `json:"involvedObjects,omitempty"`

	// **参数解释**： Job输入参数。 **取值范围**： 不涉及。
	Inputs *string `json:"inputs,omitempty"`

	// **参数解释**： Job状态。 **取值范围**： 可选值如下： - Running：任务正在运行中。 - Success：任务执行成功。 - Failed：任务执行失败。
	Phase *string `json:"phase,omitempty"`

	// **参数解释**： Job是否被挂起。 **取值范围**： 不涉及。
	Suspend *bool `json:"suspend,omitempty"`

	// **参数解释**： Job类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： Job的执行过程信息。 **取值范围**： 不涉及。
	Conditions *string `json:"conditions,omitempty"`

	// **参数解释**： Job执行失败时返回执行信息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`
}

func (o PoolJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolJob struct{}"
	}

	return strings.Join([]string{"PoolJob", string(data)}, " ")
}
