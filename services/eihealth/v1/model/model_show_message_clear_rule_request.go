package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMessageClearRuleRequest Request Object
type ShowMessageClearRuleRequest struct {
}

func (o ShowMessageClearRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessageClearRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowMessageClearRuleRequest", string(data)}, " ")
}
