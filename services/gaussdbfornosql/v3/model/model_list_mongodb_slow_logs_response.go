package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMongodbSlowLogsResponse Response Object
type ListMongodbSlowLogsResponse struct {

	// 慢日志具体信息。
	SlowLogs       *[]MongodbSlowLogDetail `json:"slow_logs,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListMongodbSlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMongodbSlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListMongodbSlowLogsResponse", string(data)}, " ")
}
