package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReportSettingRequest Request Object
type UpdateReportSettingRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 报告配置ID
	SettingId string `json:"setting_id"`

	Body *UpdateReportSettingBody `json:"body,omitempty"`
}

func (o UpdateReportSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReportSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateReportSettingRequest", string(data)}, " ")
}
