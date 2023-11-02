package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleMetricsRequest Request Object
type ListAlertRuleMetricsRequest struct {

	// 工作空间 ID。Workspace ID.
	WorkspaceId string `json:"workspace_id"`
}

func (o ListAlertRuleMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListAlertRuleMetricsRequest", string(data)}, " ")
}
