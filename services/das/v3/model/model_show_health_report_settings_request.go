package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHealthReportSettingsRequest Request Object
type ShowHealthReportSettingsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowHealthReportSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthReportSettingsRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthReportSettingsRequest", string(data)}, " ")
}
