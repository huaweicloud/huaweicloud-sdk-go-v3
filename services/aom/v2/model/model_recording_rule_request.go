package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecordingRuleRequest struct {

	// recordingRule
	RecordingRule string `json:"recording_rule"`
}

func (o RecordingRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordingRuleRequest struct{}"
	}

	return strings.Join([]string{"RecordingRuleRequest", string(data)}, " ")
}
