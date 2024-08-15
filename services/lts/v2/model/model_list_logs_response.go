package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogsResponse Response Object
type ListLogsResponse struct {

	// 日志条数。
	Count *int32 `json:"count,omitempty"`

	// 日志信息。
	Logs *[]LogContents `json:"logs,omitempty"`

	// 是否查询完成。
	IsQueryComplete *bool `json:"isQueryComplete,omitempty"`

	// 分析日志返回响应体
	AnalysisLogs   *[]map[string]string `json:"analysisLogs,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLogsResponse", string(data)}, " ")
}
