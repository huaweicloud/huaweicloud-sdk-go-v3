package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 需要设置SMN的任务信息。
type SelectedSetAlarmTaskReq struct {
	// 任务ID

	JobId string `json:"job_id"`
	// 任务状态

	Status string `json:"status"`
	// 引擎类型

	EngineType string `json:"engine_type"`
}

func (o SelectedSetAlarmTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectedSetAlarmTaskReq struct{}"
	}

	return strings.Join([]string{"SelectedSetAlarmTaskReq", string(data)}, " ")
}
