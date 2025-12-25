package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteRocketMqMigrationTaskResponse Response Object
type BatchDeleteRocketMqMigrationTaskResponse struct {

	// **参数解释**： 删除成功的任务列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及 **默认取值**： 不涉及。
	SuccessTaskList *[]string `json:"success_task_list,omitempty"`
	HttpStatusCode  int       `json:"-"`
}

func (o BatchDeleteRocketMqMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRocketMqMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteRocketMqMigrationTaskResponse", string(data)}, " ")
}
