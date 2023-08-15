package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHealthCheckResponse Response Object
type CreateHealthCheckResponse struct {
	HealthCheck *HealthCheckDetail `json:"health_check,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateHealthCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthCheckResponse struct{}"
	}

	return strings.Join([]string{"CreateHealthCheckResponse", string(data)}, " ")
}
