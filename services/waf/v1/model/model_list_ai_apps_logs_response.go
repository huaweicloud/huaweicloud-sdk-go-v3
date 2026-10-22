package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiAppsLogsResponse Response Object
type ListAiAppsLogsResponse struct {

	// 总记录数
	Total *int32 `json:"total,omitempty"`

	// 日志列表
	Items *[]LogItem `json:"items,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAiAppsLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiAppsLogsResponse struct{}"
	}

	return strings.Join([]string{"ListAiAppsLogsResponse", string(data)}, " ")
}
