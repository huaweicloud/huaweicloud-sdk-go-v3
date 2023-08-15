package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobMonitorInfoRespPayload 作业审计日志信息。
type ShowJobMonitorInfoRespPayload struct {

	// 所有作业监控信息。
	Jobs *[]ShowJobMonitorInfoRespPayloadJobs `json:"jobs,omitempty"`
}

func (o ShowJobMonitorInfoRespPayload) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobMonitorInfoRespPayload struct{}"
	}

	return strings.Join([]string{"ShowJobMonitorInfoRespPayload", string(data)}, " ")
}
