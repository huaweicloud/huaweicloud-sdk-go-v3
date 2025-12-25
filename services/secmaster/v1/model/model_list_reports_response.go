package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportsResponse Response Object
type ListReportsResponse struct {
	Body *[]ReportInfo `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportsResponse struct{}"
	}

	return strings.Join([]string{"ListReportsResponse", string(data)}, " ")
}
