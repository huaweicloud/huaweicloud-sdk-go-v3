package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmLevelResp **参数解释**： 告警级别。 **取值范围**： 取值为1、2、3、4 - 1：紧急 - 2：重要 - 3：次要 - 4：提示
type AlarmLevelResp struct {
}

func (o AlarmLevelResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmLevelResp struct{}"
	}

	return strings.Join([]string{"AlarmLevelResp", string(data)}, " ")
}
