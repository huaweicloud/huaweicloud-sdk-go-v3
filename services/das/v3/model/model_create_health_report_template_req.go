package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateHealthReportTemplateReq struct {

	// 模板名称。
	TemplateName string `json:"template_name"`

	// 模板描述。
	Description *string `json:"description,omitempty"`

	// 章节配置，key为章节编码，value为章节配置对象。
	ChaptersConfig map[string]ChapterConfigReq `json:"chapters_config"`
}

func (o CreateHealthReportTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthReportTemplateReq struct{}"
	}

	return strings.Join([]string{"CreateHealthReportTemplateReq", string(data)}, " ")
}
