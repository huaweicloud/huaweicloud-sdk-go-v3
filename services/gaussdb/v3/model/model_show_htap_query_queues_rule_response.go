package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHtapQueryQueuesRuleResponse Response Object
type ShowHtapQueryQueuesRuleResponse struct {
	QueryQueueRule *HtapQueryQueueRule `json:"query_queue_rule,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowHtapQueryQueuesRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHtapQueryQueuesRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowHtapQueryQueuesRuleResponse", string(data)}, " ")
}
