package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthReportSettingsRequest Request Object
type UpdateHealthReportSettingsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateHealthReportSettingsRequestBody `json:"body,omitempty"`
}

func (o UpdateHealthReportSettingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthReportSettingsRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthReportSettingsRequest", string(data)}, " ")
}
