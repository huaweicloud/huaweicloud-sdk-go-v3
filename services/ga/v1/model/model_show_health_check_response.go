package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowHealthCheckResponse struct {
	HealthCheck *HealthCheckDetail `json:"health_check,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHealthCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthCheckResponse struct{}"
	}

	return strings.Join([]string{"ShowHealthCheckResponse", string(data)}, " ")
}
