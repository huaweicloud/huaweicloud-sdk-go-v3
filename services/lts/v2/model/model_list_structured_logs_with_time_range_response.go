package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStructuredLogsWithTimeRangeResponse struct {
	// 此参数在请求实体中，采用json字符串格式。

	Body           map[string][]interface{} `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListStructuredLogsWithTimeRangeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStructuredLogsWithTimeRangeResponse struct{}"
	}

	return strings.Join([]string{"ListStructuredLogsWithTimeRangeResponse", string(data)}, " ")
}
