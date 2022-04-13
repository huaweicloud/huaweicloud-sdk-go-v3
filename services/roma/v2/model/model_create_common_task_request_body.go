package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务信息
type CreateCommonTaskRequestBody struct {
	Task *TaskBasicRequestBody `json:"task,omitempty"`
	// 参数类型为string，参数结构参照附录中“数据集成参数说明>RawFormDataRequest”章节

	TaskDetail *string `json:"task_detail,omitempty"`
}

func (o CreateCommonTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommonTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCommonTaskRequestBody", string(data)}, " ")
}
