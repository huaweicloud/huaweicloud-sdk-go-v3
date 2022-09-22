package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UnlockTargetEcsRequest struct {

	// 指定任务的ID
	TaskId string `json:"task_id"`
}

func (o UnlockTargetEcsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockTargetEcsRequest struct{}"
	}

	return strings.Join([]string{"UnlockTargetEcsRequest", string(data)}, " ")
}
