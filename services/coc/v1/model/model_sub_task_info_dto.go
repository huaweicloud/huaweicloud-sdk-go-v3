package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubTaskInfoDto 应急预案内部包含的作业、脚本信息
type SubTaskInfoDto struct {

	// 子任务类型
	SubTaskType *string `json:"sub_task_type,omitempty"`

	// 作业、脚本的id
	SubAssociatedTaskId *string `json:"sub_associated_task_id,omitempty"`

	// 作业、脚本的名称
	SubAssociatedTaskName *string `json:"sub_associated_task_name,omitempty"`

	// 作业、脚本的分类：自定义、公共
	SubAssociatedTaskType *string `json:"sub_associated_task_type,omitempty"`
}

func (o SubTaskInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskInfoDto struct{}"
	}

	return strings.Join([]string{"SubTaskInfoDto", string(data)}, " ")
}
