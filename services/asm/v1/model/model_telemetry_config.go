package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TelemetryConfig struct {
	Metrics *Metric `json:"metrics,omitempty"`

	AccessLogging *AccessLogging `json:"accessLogging,omitempty"`

	Tracing *Tracing `json:"tracing,omitempty"`
}

func (o TelemetryConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TelemetryConfig struct{}"
	}

	return strings.Join([]string{"TelemetryConfig", string(data)}, " ")
}
