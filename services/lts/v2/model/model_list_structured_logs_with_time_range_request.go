package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListStructuredLogsWithTimeRangeRequest struct {
	// 日志流id，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。

	LogStreamId string `json:"log_stream_id"`

	Body *QueryLtsStructLogParamsNew `json:"body,omitempty"`
}

func (o ListStructuredLogsWithTimeRangeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStructuredLogsWithTimeRangeRequest struct{}"
	}

	return strings.Join([]string{"ListStructuredLogsWithTimeRangeRequest", string(data)}, " ")
}
