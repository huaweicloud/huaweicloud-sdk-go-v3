package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportProfileResponse Response Object
type CreateReportProfileResponse struct {
	Data           *ReportProfileRespData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateReportProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportProfileResponse struct{}"
	}

	return strings.Join([]string{"CreateReportProfileResponse", string(data)}, " ")
}
