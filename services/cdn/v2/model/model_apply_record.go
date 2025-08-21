package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyRecord **参数解释：** 域名模板应用记录 **约束限制：** 不涉及
type ApplyRecord struct {

	// **参数解释：** 应用模板结果 **约束限制：** 不涉及 **取值范围：** - success: 应用模板成功 - fail: 应用模板失败 **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释：** 操作ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	OperatorId *string `json:"operator_id,omitempty"`

	// **参数解释：** 域名模板名称 **约束限制：** 不涉及 **取值范围：** - 1-100个字符 - 仅支持字母、数字、中文、下划线（_）、中横线（-） **默认取值：** 不涉及
	TmlName *string `json:"tml_name,omitempty"`

	// **参数解释：** 域名模板描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Remark *string `json:"remark,omitempty"`

	// **参数解释：** 域名模板ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TmlId *string `json:"tml_id,omitempty"`

	// **参数解释：** 域名模板应用时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ApplyTime *int64 `json:"apply_time,omitempty"`

	// **参数解释：** 域名模板类型 **约束限制：** 不涉及 **取值范围：** - 1: 系统预置模板 - 2: 租户自定义模板 **默认取值：** 不涉及
	Type *int32 `json:"type,omitempty"`

	// **参数解释：** 账户ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AccountId *string `json:"account_id,omitempty"`

	Resources *[]Resource `json:"resources,omitempty"`

	Configs *TemplateConfigs `json:"configs,omitempty"`
}

func (o ApplyRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyRecord struct{}"
	}

	return strings.Join([]string{"ApplyRecord", string(data)}, " ")
}
