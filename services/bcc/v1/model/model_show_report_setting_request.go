package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportSettingRequest Request Object
type ShowReportSettingRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 报告配置ID
	SettingId string `json:"setting_id"`
}

func (o ShowReportSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowReportSettingRequest", string(data)}, " ")
}
