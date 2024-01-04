package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFactoryTaskCompletionResYesterday struct {

	// 整时的时间点
	RecordTime *int64 `json:"record_time,omitempty"`

	// 到当前时间点完成的任务数量
	TaskCompletionNum *int64 `json:"task_completion_num,omitempty"`
}

func (o ListFactoryTaskCompletionResYesterday) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryTaskCompletionResYesterday struct{}"
	}

	return strings.Join([]string{"ListFactoryTaskCompletionResYesterday", string(data)}, " ")
}
