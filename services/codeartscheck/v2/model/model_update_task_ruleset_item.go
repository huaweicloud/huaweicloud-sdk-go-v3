package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTaskRulesetItem struct {

	// **参数解释：** 检查语言，支持的语言请关注服务更新特性文档。 **约束限制:** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Language string `json:"language"`

	// **参数解释：** 规则集ID。 **约束限制:** 长度为[1~32]个字符。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RuleSetId string `json:"rule_set_id"`

	// **参数解释：** 是否启用。 **约束限制:** 不涉及。 **取值范围：** - 0：不启用。 - 1：启用。 **默认取值：** 不涉及。
	IfUse string `json:"if_use"`

	// **参数解释：** 新/老数据。 **约束限制:** 不涉及。 **取值范围：** 1 **默认取值：** 不涉及。
	Status string `json:"status"`
}

func (o UpdateTaskRulesetItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaskRulesetItem struct{}"
	}

	return strings.Join([]string{"UpdateTaskRulesetItem", string(data)}, " ")
}
