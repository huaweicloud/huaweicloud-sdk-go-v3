package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportSettingResponse Response Object
type ShowReportSettingResponse struct {
	ReportSetting  *ReportSettingEntity `json:"report_setting,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowReportSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowReportSettingResponse", string(data)}, " ")
}
