package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportEntity 报告实体
type ReportEntity struct {

	// 报告ID
	ReportId string `json:"report_id"`

	// 报告名称
	ReportName string `json:"report_name"`

	// 报告生成时间
	ReportGeneratedTime string `json:"report_generated_time"`

	// 报告配置ID
	ReportSettingId string `json:"report_setting_id"`

	// 报告配置名称
	ReportSettingName string `json:"report_setting_name"`

	// 报告模板名称
	ReportTemplate string `json:"report_template"`

	// 报告模板ID
	ReportTemplateId string `json:"report_template_id"`
}

func (o ReportEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportEntity struct{}"
	}

	return strings.Join([]string{"ReportEntity", string(data)}, " ")
}
