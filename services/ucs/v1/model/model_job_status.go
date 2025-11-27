package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobStatus struct {

	// Job phase
	Phase *string `json:"phase,omitempty"`

	// Job reason
	Reason *string `json:"reason,omitempty"`

	// Job完成时间
	CompletionTime *string `json:"completionTime,omitempty"`

	// Job开始时间
	StartTime *string `json:"startTime,omitempty"`
}

func (o JobStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobStatus struct{}"
	}

	return strings.Join([]string{"JobStatus", string(data)}, " ")
}
