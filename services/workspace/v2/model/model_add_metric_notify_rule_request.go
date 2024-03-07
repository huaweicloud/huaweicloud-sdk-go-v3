package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddMetricNotifyRuleRequest Request Object
type AddMetricNotifyRuleRequest struct {
	Body *AddMetricNotifyRuleReq `json:"body,omitempty"`
}

func (o AddMetricNotifyRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMetricNotifyRuleRequest struct{}"
	}

	return strings.Join([]string{"AddMetricNotifyRuleRequest", string(data)}, " ")
}
