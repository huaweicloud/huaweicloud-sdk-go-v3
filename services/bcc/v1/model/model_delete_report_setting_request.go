package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReportSettingRequest Request Object
type DeleteReportSettingRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 报告配置ID
	SettingId string `json:"setting_id"`
}

func (o DeleteReportSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReportSettingRequest struct{}"
	}

	return strings.Join([]string{"DeleteReportSettingRequest", string(data)}, " ")
}
