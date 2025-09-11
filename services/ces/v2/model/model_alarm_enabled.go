package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmEnabled **参数解释**： 是否开启告警规则。     **约束限制**： 不涉及。 **取值范围**： 布尔值。 - true:开启。 - false:关闭。 **默认取值**： true
type AlarmEnabled struct {
}

func (o AlarmEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmEnabled struct{}"
	}

	return strings.Join([]string{"AlarmEnabled", string(data)}, " ")
}
