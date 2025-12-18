package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotFeatureGatesRequest Request Object
type ShowAutopilotFeatureGatesRequest struct {
}

func (o ShowAutopilotFeatureGatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotFeatureGatesRequest struct{}"
	}

	return strings.Join([]string{"ShowAutopilotFeatureGatesRequest", string(data)}, " ")
}
