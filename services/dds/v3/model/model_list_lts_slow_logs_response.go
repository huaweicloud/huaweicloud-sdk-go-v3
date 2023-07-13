package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsSlowLogsResponse Response Object
type ListLtsSlowLogsResponse struct {

	// 慢日志具体信息。
	SlowLogs       *[]SlowLogDetail `json:"slow_logs,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListLtsSlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsSlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsSlowLogsResponse", string(data)}, " ")
}
