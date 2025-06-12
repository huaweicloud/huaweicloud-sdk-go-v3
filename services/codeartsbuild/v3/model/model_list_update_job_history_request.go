package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpdateJobHistoryRequest Request Object
type ListUpdateJobHistoryRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 分页页码， 表示从此页开始查询， page_no大于等于1
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页显示的条目数量，page_size小于等于100
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListUpdateJobHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpdateJobHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListUpdateJobHistoryRequest", string(data)}, " ")
}
