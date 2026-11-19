package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendingPeriodRes **参数解释**: 报告发送的时间段 **取值范围**: - morning：0点到6点 - noon：6点到12点 - afternoon：12点到18点 - evening：18点到24点
type SendingPeriodRes struct {
}

func (o SendingPeriodRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendingPeriodRes struct{}"
	}

	return strings.Join([]string{"SendingPeriodRes", string(data)}, " ")
}
