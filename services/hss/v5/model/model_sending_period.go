package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendingPeriod **参数解释**: 报告发送的时间段 **取值范围**:   - morning：代表0点到6点   - noon：代表6点到12点   - afternoon：代表12点到18点   - evening：代表18点到24点
type SendingPeriod struct {
}

func (o SendingPeriod) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendingPeriod struct{}"
	}

	return strings.Join([]string{"SendingPeriod", string(data)}, " ")
}
