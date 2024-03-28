package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAsyncTaskResultRequest Request Object
type ShowAsyncTaskResultRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 异步任务ID
	TaskId string `json:"task_id"`
}

func (o ShowAsyncTaskResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAsyncTaskResultRequest struct{}"
	}

	return strings.Join([]string{"ShowAsyncTaskResultRequest", string(data)}, " ")
}
