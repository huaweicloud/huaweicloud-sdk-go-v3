package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopJobActionInfo 结束任务请求体。
type StopJobActionInfo struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 强制结束任务时取值true，默认false。
	IsForceStop *bool `json:"is_force_stop,omitempty"`
}

func (o StopJobActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobActionInfo struct{}"
	}

	return strings.Join([]string{"StopJobActionInfo", string(data)}, " ")
}
