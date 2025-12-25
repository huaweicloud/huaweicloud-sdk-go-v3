package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleTemplateMetricsRequest Request Object
type ListAlertRuleTemplateMetricsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o ListAlertRuleTemplateMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplateMetricsRequest struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplateMetricsRequest", string(data)}, " ")
}
