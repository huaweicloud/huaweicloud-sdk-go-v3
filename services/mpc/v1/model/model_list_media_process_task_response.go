package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMediaProcessTaskResponse struct {

	// 任务列表
	TaskArray *[]MediaProcessTaskInfo `json:"task_array,omitempty" xml:"task_array"`

	// 是否截断
	IsTruncated *int32 `json:"is_truncated,omitempty" xml:"is_truncated"`

	// 任务总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMediaProcessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMediaProcessTaskResponse struct{}"
	}

	return strings.Join([]string{"ListMediaProcessTaskResponse", string(data)}, " ")
}
