package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleSetV2 struct {

	// **参数解释：** 规则集ID。 **约束限制:** 长度为[1~32]个字符。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RulesetId *string `json:"ruleset_id,omitempty"`

	// **参数解释：** 检查语言，支持的语言请关注服务更新特性文档。 **约束限制:** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Language string `json:"language"`
}

func (o RuleSetV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleSetV2 struct{}"
	}

	return strings.Join([]string{"RuleSetV2", string(data)}, " ")
}
