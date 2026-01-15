package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditTopicReportSchedulerConfigResponse Response Object
type ShowAuditTopicReportSchedulerConfigResponse struct {
	Scheduler *SchedulerConfigBase `json:"scheduler,omitempty"`

	// 是否支持订阅  - true：支持  - false：不支持
	SmnEffective   *bool `json:"smn_effective,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowAuditTopicReportSchedulerConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditTopicReportSchedulerConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditTopicReportSchedulerConfigResponse", string(data)}, " ")
}
