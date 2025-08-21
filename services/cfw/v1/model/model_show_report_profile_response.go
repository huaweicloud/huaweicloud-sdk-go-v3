package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportProfileResponse Response Object
type ShowReportProfileResponse struct {
	Data           *ReportProfileInfoVo `json:"data,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowReportProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportProfileResponse struct{}"
	}

	return strings.Join([]string{"ShowReportProfileResponse", string(data)}, " ")
}
