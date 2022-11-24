package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警规则关联告警模板ID，如果传了，告警规则关联的策略会和告警模板策略联动变化
type AlarmTemplateId struct {
}

func (o AlarmTemplateId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplateId struct{}"
	}

	return strings.Join([]string{"AlarmTemplateId", string(data)}, " ")
}
