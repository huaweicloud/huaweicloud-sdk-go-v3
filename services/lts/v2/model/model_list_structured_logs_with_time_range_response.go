package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListStructuredLogsWithTimeRangeResponse struct {
	// 查询结构化日志结果信息。此处仅为示例，具体参数名称取决于查询的字段。

	Context        *[]string `json:"context,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListStructuredLogsWithTimeRangeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListStructuredLogsWithTimeRangeResponse struct{}"
	}

	return strings.Join([]string{"ListStructuredLogsWithTimeRangeResponse", string(data)}, " ")
}
