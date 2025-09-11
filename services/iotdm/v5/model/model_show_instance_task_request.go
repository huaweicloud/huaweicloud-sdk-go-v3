package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTaskRequest Request Object
type ShowInstanceTaskRequest struct {

	// **参数说明**：实例ID。 **取值范围**：长度不超过36，由小写字母[a-f]、数字、连接符（-）的组成。
	InstanceId string `json:"instance_id"`

	// 任务Id
	TaskId string `json:"task_id"`
}

func (o ShowInstanceTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceTaskRequest", string(data)}, " ")
}
