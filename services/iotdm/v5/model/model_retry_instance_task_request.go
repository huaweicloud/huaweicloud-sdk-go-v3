package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryInstanceTaskRequest Request Object
type RetryInstanceTaskRequest struct {

	// **参数说明**：实例ID。 **取值范围**：长度不超过36，由小写字母[a-f]、数字、连接符（-）的组成。
	InstanceId string `json:"instance_id"`

	// 任务Id
	TaskId string `json:"task_id"`
}

func (o RetryInstanceTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryInstanceTaskRequest struct{}"
	}

	return strings.Join([]string{"RetryInstanceTaskRequest", string(data)}, " ")
}
