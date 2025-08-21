package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppliedTemplateRecordRequest Request Object
type ShowAppliedTemplateRecordRequest struct {

	// **参数解释：** 域名模板ID，可以通过查询域名模板列表接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TmlId *string `json:"tml_id,omitempty"`

	// **参数解释：** 域名模板名称 **约束限制：** 不涉及 **取值范围：** - 1-100个字符 - 仅支持字母、数字、中文、下划线（_）、中横线（-） **默认取值：** 不涉及
	TmlName *string `json:"tml_name,omitempty"`

	// **参数解释：** 域名模板操作ID，可以通过本接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	OperatorId *string `json:"operator_id,omitempty"`

	// **参数解释：** 查询的页码 **约束限制：** 不涉及 **取值范围：** 0-65535 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 每页应用记录的数量 **约束限制：** 不涉及 **取值范围：** 1-10000 **默认取值：** 30
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowAppliedTemplateRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppliedTemplateRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowAppliedTemplateRecordRequest", string(data)}, " ")
}
