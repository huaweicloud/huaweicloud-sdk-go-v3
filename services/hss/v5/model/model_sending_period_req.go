package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendingPeriodReq **参数解释**: 报告发送的时间段 **约束限制**: 不涉及 **取值范围**: - morning：0点到6点 - noon：6点到12点 - afternoon：12点到18点 - evening：18点到24点  **默认取值**: 不涉及
type SendingPeriodReq struct {
}

func (o SendingPeriodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendingPeriodReq struct{}"
	}

	return strings.Join([]string{"SendingPeriodReq", string(data)}, " ")
}
