package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHealthReportTemplateResponse Response Object
type ShowHealthReportTemplateResponse struct {

	// 模板ID。
	TemplateId *int64 `json:"template_id,omitempty"`

	// 模板名称。
	TemplateName *string `json:"template_name,omitempty"`

	// 模板描述。
	Description *string `json:"description,omitempty"`

	// 是否为系统预设模板（true-系统模板 false-自定义模板）。
	System *bool `json:"system,omitempty"`

	// 章节配置，key为章节编码，value为章节配置对象。
	ChaptersConfig map[string]ChapterConfig `json:"chapters_config,omitempty"`

	// 创建时间（Unix timestamp），单位：毫秒。
	CreateAt *int64 `json:"create_at,omitempty"`

	// 更新时间（Unix timestamp），单位：毫秒。
	UpdateAt       *int64 `json:"update_at,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHealthReportTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthReportTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthReportTemplateResponse", string(data)}, " ")
}
