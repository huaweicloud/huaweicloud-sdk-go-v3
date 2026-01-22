package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRocketMqMigrationTaskResponse Response Object
type CreateRocketMqMigrationTaskResponse struct {

	// **参数解释**： 任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRocketMqMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRocketMqMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateRocketMqMigrationTaskResponse", string(data)}, " ")
}
