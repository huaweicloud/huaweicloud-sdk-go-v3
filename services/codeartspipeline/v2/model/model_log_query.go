package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogQuery 流水线日志查询请求体 startOffset 和 endOffset 均设置为 0，则代表查询全量日志。
type LogQuery struct {

	// 日志起始偏移
	StartOffset *int64 `json:"start_offset,omitempty"`

	// 日志结束偏移
	EndOffset *int64 `json:"end_offset,omitempty"`

	// 最大日志行数
	Limit int64 `json:"limit"`

	// 排序规则[\"asc\",\"desc\"]
	Sort string `json:"sort"`

	Level *string `json:"level,omitempty"`

	JobRunId *string `json:"job_run_id,omitempty"`
}

func (o LogQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogQuery struct{}"
	}

	return strings.Join([]string{"LogQuery", string(data)}, " ")
}
