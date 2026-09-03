package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopBinlogTaskRequestBody 停止binlog解析任务请求体
type StopBinlogTaskRequestBody struct {

	// 任务ID
	TaskId int64 `json:"task_id"`
}

func (o StopBinlogTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopBinlogTaskRequestBody struct{}"
	}

	return strings.Join([]string{"StopBinlogTaskRequestBody", string(data)}, " ")
}
