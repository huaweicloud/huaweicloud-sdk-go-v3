package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleMetricsRequest Request Object
type ListAlertRuleMetricsRequest struct {

	// workspace_id
	WorkspaceId string `json:"workspace_id"`
}

func (o ListAlertRuleMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListAlertRuleMetricsRequest", string(data)}, " ")
}
