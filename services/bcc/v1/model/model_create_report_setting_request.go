package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReportSettingRequest Request Object
type CreateReportSettingRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	Body *ReportSetting `json:"body,omitempty"`
}

func (o CreateReportSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReportSettingRequest struct{}"
	}

	return strings.Join([]string{"CreateReportSettingRequest", string(data)}, " ")
}
