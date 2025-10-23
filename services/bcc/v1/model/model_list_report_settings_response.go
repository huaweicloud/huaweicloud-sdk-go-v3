package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportSettingsResponse Response Object
type ListReportSettingsResponse struct {

	// 报告配置总条数
	Count *int64 `json:"count,omitempty"`

	// 配置列表
	ReportSettings *[]ReportSettingEntity `json:"report_settings,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListReportSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportSettingsResponse struct{}"
	}

	return strings.Join([]string{"ListReportSettingsResponse", string(data)}, " ")
}
