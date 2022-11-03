package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteHealthCheckRequest struct {

	// 健康检查ID。
	HealthCheckId string `json:"health_check_id"`
}

func (o DeleteHealthCheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHealthCheckRequest struct{}"
	}

	return strings.Join([]string{"DeleteHealthCheckRequest", string(data)}, " ")
}
