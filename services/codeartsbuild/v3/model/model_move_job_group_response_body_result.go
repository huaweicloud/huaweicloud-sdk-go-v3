package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveJobGroupResponseBodyResult 任务组
type MoveJobGroupResponseBodyResult struct {

	// 任务编号
	JobId *string `json:"job_id,omitempty"`

	// 分组路径
	GroupPathId *string `json:"group_path_id,omitempty"`
}

func (o MoveJobGroupResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveJobGroupResponseBodyResult struct{}"
	}

	return strings.Join([]string{"MoveJobGroupResponseBodyResult", string(data)}, " ")
}
