package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmStatus **参数解释**： 处理方式 **取值范围**: - confirmed：已确认 - unconfirmed：未确认
type ConfirmStatus struct {
}

func (o ConfirmStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmStatus struct{}"
	}

	return strings.Join([]string{"ConfirmStatus", string(data)}, " ")
}
