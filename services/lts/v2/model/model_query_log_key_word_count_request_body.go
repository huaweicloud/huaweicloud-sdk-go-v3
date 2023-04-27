package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志关键词统计请求体。
type QueryLogKeyWordCountRequestBody struct {

	// 开始时间
	StartTime string `json:"start_time"`

	// 结束时间
	EndTime string `json:"end_time"`

	// 步长间隔
	StepInterval int64 `json:"step_interval"`

	// 日志组ID
	GroupId string `json:"group_id"`

	// 日志流ID
	StreamId string `json:"stream_id"`

	// 关键词
	KeyWord string `json:"key_word"`

	// 日志迭代查询，默认为false（不开启迭代），true为开启迭代。
	IsIterative *bool `json:"is_iterative,omitempty"`
}

func (o QueryLogKeyWordCountRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryLogKeyWordCountRequestBody struct{}"
	}

	return strings.Join([]string{"QueryLogKeyWordCountRequestBody", string(data)}, " ")
}
