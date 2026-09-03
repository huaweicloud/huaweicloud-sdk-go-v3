package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceHealthReportTaskNewRequestBody 创建实例健康报告任务请求体
type CreateInstanceHealthReportTaskNewRequestBody struct {

	// 开始时间（Unix时间戳，毫秒）
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix时间戳，毫秒）
	EndAt int64 `json:"end_at"`
}

func (o CreateInstanceHealthReportTaskNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceHealthReportTaskNewRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceHealthReportTaskNewRequestBody", string(data)}, " ")
}
