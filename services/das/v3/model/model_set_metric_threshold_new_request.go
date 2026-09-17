package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMetricThresholdNewRequest Request Object
type SetMetricThresholdNewRequest struct {
	Body *SetMetricThresholdNewRequestBody `json:"body,omitempty"`
}

func (o SetMetricThresholdNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMetricThresholdNewRequest struct{}"
	}

	return strings.Join([]string{"SetMetricThresholdNewRequest", string(data)}, " ")
}
