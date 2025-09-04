package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveJobGroupRequestBody 创建更新分组返回体
type MoveJobGroupRequestBody struct {

	// 任务分组id
	GroupId *string `json:"group_id,omitempty"`

	// 任务组
	Jobs *[]MoveJobGroupRequestBodyJobs `json:"jobs,omitempty"`
}

func (o MoveJobGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveJobGroupRequestBody struct{}"
	}

	return strings.Join([]string{"MoveJobGroupRequestBody", string(data)}, " ")
}
