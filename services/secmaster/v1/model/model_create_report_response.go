package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportResponse Response Object
type CreateReportResponse struct {

	// 响应示例
	Body *string `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportResponse struct{}"
	}

	return strings.Join([]string{"CreateReportResponse", string(data)}, " ")
}
