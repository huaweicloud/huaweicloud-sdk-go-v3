package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetMetricThresholdNewResponse Response Object
type SetMetricThresholdNewResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SetMetricThresholdNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetMetricThresholdNewResponse struct{}"
	}

	return strings.Join([]string{"SetMetricThresholdNewResponse", string(data)}, " ")
}
