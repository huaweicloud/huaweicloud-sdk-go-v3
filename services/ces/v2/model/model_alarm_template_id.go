package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmTemplateId 告警规则关联告警模板ID
type AlarmTemplateId struct {
}

func (o AlarmTemplateId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplateId struct{}"
	}

	return strings.Join([]string{"AlarmTemplateId", string(data)}, " ")
}
