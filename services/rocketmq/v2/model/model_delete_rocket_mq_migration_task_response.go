package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRocketMqMigrationTaskResponse Response Object
type DeleteRocketMqMigrationTaskResponse struct {

	// **参数解释**： 删除成功的任务列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及 **默认取值**： 不涉及。
	SuccessTaskList *[]string `json:"success_task_list,omitempty"`
	HttpStatusCode  int       `json:"-"`
}

func (o DeleteRocketMqMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRocketMqMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteRocketMqMigrationTaskResponse", string(data)}, " ")
}
