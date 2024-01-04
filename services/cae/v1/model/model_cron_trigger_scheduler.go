package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CronTriggerScheduler struct {

	// 触发时间点，以五位cron表达式表示。
	Cron *string `json:"cron,omitempty"`

	// 要求达到的实例数。
	TargetReplica *int32 `json:"target_replica,omitempty"`
}

func (o CronTriggerScheduler) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CronTriggerScheduler struct{}"
	}

	return strings.Join([]string{"CronTriggerScheduler", string(data)}, " ")
}
