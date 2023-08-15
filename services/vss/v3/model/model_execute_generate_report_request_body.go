package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteGenerateReportRequestBody struct {

	// 任务ID
	TaskId string `json:"task_id"`
}

func (o ExecuteGenerateReportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGenerateReportRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteGenerateReportRequestBody", string(data)}, " ")
}
