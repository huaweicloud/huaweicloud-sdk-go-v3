package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReportProfileResponse Response Object
type DeleteReportProfileResponse struct {
	Data           *ReportProfileRespData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o DeleteReportProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReportProfileResponse struct{}"
	}

	return strings.Join([]string{"DeleteReportProfileResponse", string(data)}, " ")
}
