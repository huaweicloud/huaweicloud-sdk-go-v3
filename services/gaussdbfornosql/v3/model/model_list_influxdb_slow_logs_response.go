package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInfluxdbSlowLogsResponse Response Object
type ListInfluxdbSlowLogsResponse struct {

	// Influxdb慢日志列表
	SlowLogs       *[]InfluxdbSlowLogDetail `json:"slow_logs,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListInfluxdbSlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInfluxdbSlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListInfluxdbSlowLogsResponse", string(data)}, " ")
}
