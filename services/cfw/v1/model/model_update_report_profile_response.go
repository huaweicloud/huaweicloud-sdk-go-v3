package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReportProfileResponse Response Object
type UpdateReportProfileResponse struct {
	Data           *ReportProfileRespData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateReportProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReportProfileResponse struct{}"
	}

	return strings.Join([]string{"UpdateReportProfileResponse", string(data)}, " ")
}
