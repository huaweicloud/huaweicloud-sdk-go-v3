package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHealthCheckRequest struct {

	// 健康检查ID。
	HealthCheckId string `json:"health_check_id"`

	Body *UpdateHealthCheckRequestBody `json:"body,omitempty"`
}

func (o UpdateHealthCheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthCheckRequest struct{}"
	}

	return strings.Join([]string{"UpdateHealthCheckRequest", string(data)}, " ")
}
