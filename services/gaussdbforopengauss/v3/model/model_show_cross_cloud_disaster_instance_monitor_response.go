package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCrossCloudDisasterInstanceMonitorResponse Response Object
type ShowCrossCloudDisasterInstanceMonitorResponse struct {

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 容灾状态。
	Status *string `json:"status,omitempty"`

	// 数据恢复点目标。
	Rpo *string `json:"rpo,omitempty"`

	// 数据恢复时间目标。
	Rto *string `json:"rto,omitempty"`

	// rpo阈值。
	RpoThreshold *string `json:"rpo_threshold,omitempty"`

	// rto阈值。
	RtoThreshold *string `json:"rto_threshold,omitempty"`

	// 主从切换进度。该值为一个百分数。例如：40%。
	SwitchoverProgress *string `json:"switchover_progress,omitempty"`

	// 容灾升主进度。该值为一个百分数。例如：40%。
	FailoverProgress *string `json:"failover_progress,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowCrossCloudDisasterInstanceMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCrossCloudDisasterInstanceMonitorResponse struct{}"
	}

	return strings.Join([]string{"ShowCrossCloudDisasterInstanceMonitorResponse", string(data)}, " ")
}
