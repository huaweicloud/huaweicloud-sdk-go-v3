package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCassandraSlowLogsResponse Response Object
type ListCassandraSlowLogsResponse struct {

	// 慢日志具体信息。
	SlowLogs       *[]CassandraSlowLogDetail `json:"slow_logs,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListCassandraSlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCassandraSlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListCassandraSlowLogsResponse", string(data)}, " ")
}
