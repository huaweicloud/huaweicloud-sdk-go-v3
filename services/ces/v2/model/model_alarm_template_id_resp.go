package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmTemplateIdResp **参数解释**： 告警规则关联告警模板ID     **取值范围**： 以at开头，只包含字母、数字，长度为[2,64]个字符。
type AlarmTemplateIdResp struct {
}

func (o AlarmTemplateIdResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplateIdResp struct{}"
	}

	return strings.Join([]string{"AlarmTemplateIdResp", string(data)}, " ")
}
