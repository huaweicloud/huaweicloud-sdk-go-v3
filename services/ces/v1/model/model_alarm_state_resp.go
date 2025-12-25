package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmStateResp **参数解释**： 告警状态。 **取值范围**： 只能为ok、alarm、insufficient_data。 - ok：正常 - alarm：告警 - insufficient_data：数据不足
type AlarmStateResp struct {
}

func (o AlarmStateResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmStateResp struct{}"
	}

	return strings.Join([]string{"AlarmStateResp", string(data)}, " ")
}
