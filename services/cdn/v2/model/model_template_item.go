package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateItem **参数解释：** 查询模板配置 **约束限制：** 不涉及
type TemplateItem struct {

	// **参数解释：** 账户ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AccountId *string `json:"account_id,omitempty"`

	// **参数解释：** 域名模板名称 **约束限制：** 不涉及 **取值范围：** - 1-100个字符 - 仅支持字母、数字、中文、下划线（_）、中横线（-） **默认取值：** 不涉及
	TmlName *string `json:"tml_name,omitempty"`

	// **参数解释：** 域名模板描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Remark *string `json:"remark,omitempty"`

	// **参数解释：** 域名模板ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TmlId *string `json:"tml_id,omitempty"`

	// **参数解释：** 域名模板类型 **约束限制：** 不涉及 **取值范围：** - 1: 系统预置模板 - 2: 租户自定义模板 **默认取值：** 不涉及
	Type *int32 `json:"type,omitempty"`

	// **参数解释：** 创建时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释：** 修改时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ModifyTime *int64 `json:"modify_time,omitempty"`

	Configs *TemplateConfigs `json:"configs,omitempty"`
}

func (o TemplateItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateItem struct{}"
	}

	return strings.Join([]string{"TemplateItem", string(data)}, " ")
}
