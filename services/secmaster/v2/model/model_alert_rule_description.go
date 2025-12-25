package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertRuleDescription Alert rule description 告警规则描述
type AlertRuleDescription struct {
}

func (o AlertRuleDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleDescription struct{}"
	}

	return strings.Join([]string{"AlertRuleDescription", string(data)}, " ")
}
