package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExportSlowQueryLogsResponse struct {

	// 慢SQL集合。当集合为空时，说明慢SQL已全部导出。
	SlowLogs *[]SlowLog `json:"slow_logs,omitempty"`

	// 获取下一页所需的标识符。marker仅在3分钟内有效。
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportSlowQueryLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSlowQueryLogsResponse struct{}"
	}

	return strings.Join([]string{"ExportSlowQueryLogsResponse", string(data)}, " ")
}
