package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduledTaskEntity struct {

	// **参数解释**： 任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 任务名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 用户名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 用户ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 任务参数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Params *string `json:"params,omitempty"`

	// **参数解释**： 任务状态。 **约束限制**： 不涉及。 **取值范围**： - CREATED：定时任务状态为创建成功。 - SUCCESS：定时任务状态为成功。 - FAILED：定时任务状态为失败。 - DELETED：定时任务状态为已删除。 - EXECUTING：定时任务状态为执行中。 - CANCELLED：定时任务状态为取消。 **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 创建时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**： 更新时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释**： 定时任务执行时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ScheduleAt *string `json:"schedule_at,omitempty"`
}

func (o ScheduledTaskEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTaskEntity struct{}"
	}

	return strings.Join([]string{"ScheduledTaskEntity", string(data)}, " ")
}
