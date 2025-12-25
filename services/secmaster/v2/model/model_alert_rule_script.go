package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertRuleScript Script 脚本
type AlertRuleScript struct {
}

func (o AlertRuleScript) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleScript struct{}"
	}

	return strings.Join([]string{"AlertRuleScript", string(data)}, " ")
}
