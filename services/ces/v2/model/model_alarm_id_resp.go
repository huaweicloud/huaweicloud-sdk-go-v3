package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmIdResp **参数解释**： 告警规则id。如 al123232232341232132 **取值范围**： 以al开头，后跟22个数字或字母。长度为24个字符。
type AlarmIdResp struct {
}

func (o AlarmIdResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmIdResp struct{}"
	}

	return strings.Join([]string{"AlarmIdResp", string(data)}, " ")
}
