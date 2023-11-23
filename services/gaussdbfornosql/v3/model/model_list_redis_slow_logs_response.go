package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRedisSlowLogsResponse Response Object
type ListRedisSlowLogsResponse struct {

	// 慢日志具体信息。
	SlowLogs       *[]RedisSlowLogDetail `json:"slow_logs,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListRedisSlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedisSlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListRedisSlowLogsResponse", string(data)}, " ")
}
