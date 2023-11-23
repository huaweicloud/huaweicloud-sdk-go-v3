package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMongodbErrorLogsResponse Response Object
type ListMongodbErrorLogsResponse struct {

	// 错误日志具体信息。
	ErrorLogs      *[]MongodbErrorLogDetail `json:"error_logs,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListMongodbErrorLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMongodbErrorLogsResponse struct{}"
	}

	return strings.Join([]string{"ListMongodbErrorLogsResponse", string(data)}, " ")
}
