package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmLevel **参数解释**： 告警级别。     **约束限制**： 不涉及。 **取值范围**： 只能为1、2、3、4。 - 1为紧急 - 2为重要 - 3为次要 - 4为提示           **默认取值**： 不涉及。
type AlarmLevel struct {
}

func (o AlarmLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmLevel struct{}"
	}

	return strings.Join([]string{"AlarmLevel", string(data)}, " ")
}
