package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHealthCheckRequest Request Object
type ShowHealthCheckRequest struct {

	// 健康检查ID。
	HealthCheckId string `json:"health_check_id"`
}

func (o ShowHealthCheckRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHealthCheckRequest struct{}"
	}

	return strings.Join([]string{"ShowHealthCheckRequest", string(data)}, " ")
}
