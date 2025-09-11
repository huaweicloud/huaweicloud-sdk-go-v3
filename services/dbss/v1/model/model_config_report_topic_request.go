package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigReportTopicRequest struct {
	Scheduler *SchedulerBean `json:"scheduler"`
}

func (o ConfigReportTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigReportTopicRequest struct{}"
	}

	return strings.Join([]string{"ConfigReportTopicRequest", string(data)}, " ")
}
