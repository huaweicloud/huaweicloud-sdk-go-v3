package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReportResponse Response Object
type DeleteReportResponse struct {

	// 响应示例
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReportResponse struct{}"
	}

	return strings.Join([]string{"DeleteReportResponse", string(data)}, " ")
}
