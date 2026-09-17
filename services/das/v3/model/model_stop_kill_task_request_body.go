package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopKillTaskRequestBody Stop Kill Task请求体
type StopKillTaskRequestBody struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o StopKillTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopKillTaskRequestBody struct{}"
	}

	return strings.Join([]string{"StopKillTaskRequestBody", string(data)}, " ")
}
