package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricThresholdRequest Request Object
type ShowMetricThresholdRequest struct {
	Body *ShowMetricThresholdRequestBody `json:"body,omitempty"`
}

func (o ShowMetricThresholdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricThresholdRequest struct{}"
	}

	return strings.Join([]string{"ShowMetricThresholdRequest", string(data)}, " ")
}
