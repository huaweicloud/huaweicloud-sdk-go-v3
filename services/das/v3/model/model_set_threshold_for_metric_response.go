package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetThresholdForMetricResponse Response Object
type SetThresholdForMetricResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SetThresholdForMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetThresholdForMetricResponse struct{}"
	}

	return strings.Join([]string{"SetThresholdForMetricResponse", string(data)}, " ")
}
