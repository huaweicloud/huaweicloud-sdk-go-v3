package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReportResponse Response Object
type UpdateReportResponse struct {

	// 响应示例
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReportResponse struct{}"
	}

	return strings.Join([]string{"UpdateReportResponse", string(data)}, " ")
}
