package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportSettingEntity 报告配置实体
type ReportSettingEntity struct {

	// 配置ID
	SettingId string `json:"setting_id"`

	// 本配置所有已经生成的报告数量，不包含已经删除的报告
	ReportCount int32 `json:"report_count"`

	// 报告配置的创建时间
	CreateTime string `json:"create_time"`

	// 报告配置的上一次更新时间
	LastUpdateTime string `json:"last_update_time"`

	// 报告关联的模板ID
	ReportTemplateId string `json:"report_template_id"`

	SettingContent *ReportSetting `json:"setting_content"`
}

func (o ReportSettingEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportSettingEntity struct{}"
	}

	return strings.Join([]string{"ReportSettingEntity", string(data)}, " ")
}
