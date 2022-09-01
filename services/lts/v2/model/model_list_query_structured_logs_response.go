package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListQueryStructuredLogsResponse struct {

	// 日志信息。
	StructLogs     *[]StructLogContents `json:"struct_logs,omitempty" xml:"struct_logs"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListQueryStructuredLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryStructuredLogsResponse struct{}"
	}

	return strings.Join([]string{"ListQueryStructuredLogsResponse", string(data)}, " ")
}
