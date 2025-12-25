package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertRuleTemplateDescription 告警规则模板描述
type AlertRuleTemplateDescription struct {
}

func (o AlertRuleTemplateDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertRuleTemplateDescription struct{}"
	}

	return strings.Join([]string{"AlertRuleTemplateDescription", string(data)}, " ")
}
