package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditTopicReportSchedulerConfigResponse Response Object
type UpdateAuditTopicReportSchedulerConfigResponse struct {

	// 状态  - success：成功  - fail：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditTopicReportSchedulerConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditTopicReportSchedulerConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditTopicReportSchedulerConfigResponse", string(data)}, " ")
}
