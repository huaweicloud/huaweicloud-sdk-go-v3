package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHealthCheckRequestBody update Health Check request
type UpdateHealthCheckRequestBody struct {
	HealthCheck *UpdateHealthCheckOption `json:"health_check"`
}

func (o UpdateHealthCheckRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthCheckRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHealthCheckRequestBody", string(data)}, " ")
}
