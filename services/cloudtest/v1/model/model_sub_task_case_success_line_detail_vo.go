package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubTaskCaseSuccessLineDetailVo struct {

	// 统计时间
	Date *int64 `json:"date,omitempty"`

	// 成功率
	PassRate *float32 `json:"pass_rate,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`
}

func (o SubTaskCaseSuccessLineDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskCaseSuccessLineDetailVo struct{}"
	}

	return strings.Join([]string{"SubTaskCaseSuccessLineDetailVo", string(data)}, " ")
}
