package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopJobInstanceRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
	// 作业实例id.

	InstanceId string `json:"instance_id"`
}

func (o StopJobInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopJobInstanceRequest", string(data)}, " ")
}
