package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteCaptureTaskDto struct {

	// 抓包任务id列表
	TaskIds []string `json:"task_ids"`
}

func (o DeleteCaptureTaskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCaptureTaskDto struct{}"
	}

	return strings.Join([]string{"DeleteCaptureTaskDto", string(data)}, " ")
}
