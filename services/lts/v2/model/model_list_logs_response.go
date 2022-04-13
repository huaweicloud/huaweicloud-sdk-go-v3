package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLogsResponse struct {
	// 日志条数。

	Count *int32 `json:"count,omitempty"`
	// 日志信息。

	Logs           *[]LogContents `json:"logs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLogsResponse", string(data)}, " ")
}
