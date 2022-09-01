package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志关键词统计请求体。
type QueryLogKeyWordCountRequestBody struct {

	// 开始时间
	StartTime string `json:"start_time" xml:"start_time"`

	// 结束时间
	EndTime string `json:"end_time" xml:"end_time"`

	// 步长间隔
	StepInterval int64 `json:"step_interval" xml:"step_interval"`

	// 日志组ID
	GroupId string `json:"group_id" xml:"group_id"`

	// 日志流ID
	StreamId string `json:"stream_id" xml:"stream_id"`

	// 关键词
	KeyWord string `json:"key_word" xml:"key_word"`
}

func (o QueryLogKeyWordCountRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryLogKeyWordCountRequestBody struct{}"
	}

	return strings.Join([]string{"QueryLogKeyWordCountRequestBody", string(data)}, " ")
}
