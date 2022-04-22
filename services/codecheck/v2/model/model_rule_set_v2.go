package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleSetV2 struct {

	// 规则集id，需要从web界面获取
	RulesetId *string `json:"ruleset_id,omitempty"`

	// 检查语言，支持cpp,java,js,python,php,css,html,go,typescript
	Language string `json:"language"`
}

func (o RuleSetV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleSetV2 struct{}"
	}

	return strings.Join([]string{"RuleSetV2", string(data)}, " ")
}
