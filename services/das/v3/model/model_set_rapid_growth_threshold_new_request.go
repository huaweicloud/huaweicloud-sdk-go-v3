package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRapidGrowthThresholdNewRequest Request Object
type SetRapidGrowthThresholdNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SetRapidGrowthThresholdNewRequestBody `json:"body,omitempty"`
}

func (o SetRapidGrowthThresholdNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRapidGrowthThresholdNewRequest struct{}"
	}

	return strings.Join([]string{"SetRapidGrowthThresholdNewRequest", string(data)}, " ")
}
