package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthReportSettingsResponse Response Object
type UpdateHealthReportSettingsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateHealthReportSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthReportSettingsResponse struct{}"
	}

	return strings.Join([]string{"UpdateHealthReportSettingsResponse", string(data)}, " ")
}
