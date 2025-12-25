package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertRuleName Alert rule name 告警规则名称
type AlertRuleName struct {
}

func (o AlertRuleName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleName struct{}"
	}

	return strings.Join([]string{"AlertRuleName", string(data)}, " ")
}
