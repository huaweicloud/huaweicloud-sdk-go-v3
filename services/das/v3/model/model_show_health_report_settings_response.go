package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHealthReportSettingsResponse Response Object
type ShowHealthReportSettingsResponse struct {

	// 更新时间
	UpdateAt *int64 `json:"update_at,omitempty"`

	// 是否进行AI异常检测
	DoAiAnomalyDetection *bool `json:"do_ai_anomaly_detection,omitempty"`
	HttpStatusCode       int   `json:"-"`
}

func (o ShowHealthReportSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthReportSettingsResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthReportSettingsResponse", string(data)}, " ")
}
