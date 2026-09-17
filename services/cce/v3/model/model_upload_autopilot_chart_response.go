package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAutopilotChartResponse Response Object
type UploadAutopilotChartResponse struct {

	// **参数解释：** 模板ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 模板名称 **约束限制：** 最长64个字符 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 模板值 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Values *string `json:"values,omitempty"`

	// **参数解释：** 模板翻译资源 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Translate *string `json:"translate,omitempty"`

	// **参数解释：** 模板介绍 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Instruction *string `json:"instruction,omitempty"`

	// **参数解释：** 模板版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释：** 模板描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 模板的来源 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Source *string `json:"source,omitempty"`

	// **参数解释：** 模板的图标链接 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	IconUrl *string `json:"icon_url,omitempty"`

	// **参数解释：** 是否公开模板 **约束限制：** 不涉及 **取值范围：** - true：公开模板 - false：不公开模板  **默认取值：** false
	Public *bool `json:"public,omitempty"`

	// **参数解释：** 模板的链接 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChartUrl *string `json:"chart_url,omitempty"`

	// **参数解释：** 创建时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释：** 更新时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	UpdateAt       *string `json:"update_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadAutopilotChartResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAutopilotChartResponse struct{}"
	}

	return strings.Join([]string{"UploadAutopilotChartResponse", string(data)}, " ")
}
