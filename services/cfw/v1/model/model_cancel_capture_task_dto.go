package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CancelCaptureTaskDto struct {

	// 抓包任务id，可通过[查询抓包任务接口](ListCaptureTask.xml)查询获得，通过返回值中的data.records.task_id（.表示各对象之间层级的区分）获得。
	TaskId string `json:"task_id"`
}

func (o CancelCaptureTaskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelCaptureTaskDto struct{}"
	}

	return strings.Join([]string{"CancelCaptureTaskDto", string(data)}, " ")
}
