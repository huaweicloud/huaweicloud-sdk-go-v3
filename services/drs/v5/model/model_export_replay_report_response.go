package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportReplayReportResponse Response Object
type ExportReplayReportResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ExportReplayReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportReplayReportResponse struct{}"
	}

	return strings.Join([]string{"ExportReplayReportResponse", string(data)}, " ")
}
