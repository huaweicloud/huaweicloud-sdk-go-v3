package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHealthReportSettingsRequestBody struct {

	// 是否进行AI异常检测
	DoAiAnomalyDetection bool `json:"do_ai_anomaly_detection"`
}

func (o UpdateHealthReportSettingsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthReportSettingsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHealthReportSettingsRequestBody", string(data)}, " ")
}
