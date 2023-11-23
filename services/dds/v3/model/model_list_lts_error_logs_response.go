package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsErrorLogsResponse Response Object
type ListLtsErrorLogsResponse struct {

	// 错误日志具体信息。
	ErrorLogs      *[]ErrorLogDetail `json:"error_logs,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListLtsErrorLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsErrorLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsErrorLogsResponse", string(data)}, " ")
}
