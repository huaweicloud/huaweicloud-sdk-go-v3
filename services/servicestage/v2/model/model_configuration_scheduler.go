package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationScheduler 调度策略
type ConfigurationScheduler struct {
	Affinity *SchedulerAffinity `json:"affinity,omitempty"`

	AntiAffinity *SchedulerAffinity `json:"anti-affinity,omitempty"`
}

func (o ConfigurationScheduler) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationScheduler struct{}"
	}

	return strings.Join([]string{"ConfigurationScheduler", string(data)}, " ")
}
