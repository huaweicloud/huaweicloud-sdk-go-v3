package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateHealthCheckReq struct {

	// 健康检查ID。 通过云解析服务的管理控制台，在健康检查的详情页面中获取。
	HealthCheckId *string `json:"health_check_id,omitempty"`
}

func (o AssociateHealthCheckReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateHealthCheckReq struct{}"
	}

	return strings.Join([]string{"AssociateHealthCheckReq", string(data)}, " ")
}
