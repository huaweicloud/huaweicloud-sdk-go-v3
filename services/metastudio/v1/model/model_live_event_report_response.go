package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveEventReportResponse Response Object
type LiveEventReportResponse struct {

	// 刷新后的直播事件上传URL
	LiveEventReportUrl *string `json:"live_event_report_url,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o LiveEventReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveEventReportResponse struct{}"
	}

	return strings.Join([]string{"LiveEventReportResponse", string(data)}, " ")
}
