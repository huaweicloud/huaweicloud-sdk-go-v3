package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmNameResp **参数解释**： 告警名称。 **取值范围**： 只能包含0-9/a-z/A-Z/_/-或汉字，字符长度为[1,128]。
type AlarmNameResp struct {
}

func (o AlarmNameResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmNameResp struct{}"
	}

	return strings.Join([]string{"AlarmNameResp", string(data)}, " ")
}
