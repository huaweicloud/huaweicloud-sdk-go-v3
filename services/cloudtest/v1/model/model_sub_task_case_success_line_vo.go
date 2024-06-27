package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubTaskCaseSuccessLineVo struct {

	// 用例成功率统计信息
	Detail *[]SubTaskCaseSuccessLineDetailVo `json:"detail,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`
}

func (o SubTaskCaseSuccessLineVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskCaseSuccessLineVo struct{}"
	}

	return strings.Join([]string{"SubTaskCaseSuccessLineVo", string(data)}, " ")
}
