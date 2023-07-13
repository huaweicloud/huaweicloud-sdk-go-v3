package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateHealthCheckResponse Response Object
type DisassociateHealthCheckResponse struct {

	// 健康检查ID。 通过云解析服务的管理控制台，在健康检查的详情页面中获取。
	HealthCheckId  *string `json:"health_check_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisassociateHealthCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateHealthCheckResponse struct{}"
	}

	return strings.Join([]string{"DisassociateHealthCheckResponse", string(data)}, " ")
}
