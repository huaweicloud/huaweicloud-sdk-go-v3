package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthCheckResponse Response Object
type UpdateHealthCheckResponse struct {
	HealthCheck *HealthCheckDetail `json:"health_check,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHealthCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthCheckResponse struct{}"
	}

	return strings.Join([]string{"UpdateHealthCheckResponse", string(data)}, " ")
}
