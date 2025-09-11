package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmName **参数解释**： 告警名称。 **约束限制**： 不涉及。 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，长度[1，128]个字符。           **默认取值**： 不涉及。
type AlarmName struct {
}

func (o AlarmName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmName struct{}"
	}

	return strings.Join([]string{"AlarmName", string(data)}, " ")
}
