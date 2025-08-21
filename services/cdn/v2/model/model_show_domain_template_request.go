package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainTemplateRequest Request Object
type ShowDomainTemplateRequest struct {

	// **参数解释：** 域名模板名称 **约束限制：** 不涉及 **取值范围：** - 1-100个字符 - 仅支持字母、数字、中文、下划线（_）、中横线（-） **默认取值：** 不涉及
	TmlName *string `json:"tml_name,omitempty"`

	// **参数解释：** 域名模板ID，可以通过本接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TmlId *string `json:"tml_id,omitempty"`

	// **参数解释：** 域名模板类型 **约束限制：** 不涉及 **取值范围：** - 1: 系统预置模板 - 2: 租户自定义模板 **默认取值：** 不涉及
	TmlType *int32 `json:"tml_type,omitempty"`

	// **参数解释：** 分页大小 **约束限制：** 不涉及 **取值范围：** 1-10000 **默认取值：** 30
	Limit *string `json:"limit,omitempty"`

	// **参数解释：** 查询的页码 **约束限制：** 不涉及 **取值范围：** 0-65535 **默认取值：** 0
	Offset *string `json:"offset,omitempty"`
}

func (o ShowDomainTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainTemplateRequest", string(data)}, " ")
}
