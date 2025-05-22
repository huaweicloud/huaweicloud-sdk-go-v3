package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBuildInfoRecordRequest Request Object
type ListBuildInfoRecordRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 区间开始时间，格式yyyy-MM-dd HH:mm:ss。
	StartTime string `json:"start_time"`

	// 区间结束时间，格式yyyy-MM-dd HH:mm:ss。
	EndTime string `json:"end_time"`

	// 分页页码，表示从此页开始查询， page_index大于等于0
	PageIndex *int32 `json:"page_index,omitempty"`

	// 每页显示的条目数量，page_size小于等于100
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListBuildInfoRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuildInfoRecordRequest struct{}"
	}

	return strings.Join([]string{"ListBuildInfoRecordRequest", string(data)}, " ")
}
