package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobMonitorInfoRespPayloadJobs 所有作业监控信息。
type ShowJobMonitorInfoRespPayloadJobs struct {

	// 作业ID。
	JobId *int64 `json:"job_id,omitempty"`

	Metrics *ShowJobMonitorInfoRespPayloadJobsMetrics `json:"metrics,omitempty"`
}

func (o ShowJobMonitorInfoRespPayloadJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobMonitorInfoRespPayloadJobs struct{}"
	}

	return strings.Join([]string{"ShowJobMonitorInfoRespPayloadJobs", string(data)}, " ")
}
