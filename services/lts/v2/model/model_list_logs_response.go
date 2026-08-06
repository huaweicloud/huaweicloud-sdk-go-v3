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

	// **参数解释：** 分页查询时，若返回结果中包含该字段，下一次请求体需要增加scroll_Id参数参与分页查询。 **取值范围：** 不涉及。
	ScrollId *string `json:"scrollId,omitempty"`

	// 分析日志返回响应体
	AnalysisLogs   *[]interface{} `json:"analysisLogs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLogsResponse", string(data)}, " ")
}
