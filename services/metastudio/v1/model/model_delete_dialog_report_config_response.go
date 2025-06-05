package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDialogReportConfigResponse Response Object
type DeleteDialogReportConfigResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDialogReportConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDialogReportConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteDialogReportConfigResponse", string(data)}, " ")
}
