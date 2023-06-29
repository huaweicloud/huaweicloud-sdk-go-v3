package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportRoutePolicy 出方向路由策略
type ExportRoutePolicy struct {

	// 出方向Ipv4协议路由策略id
	ExportPolicyId *string `json:"export_policy_id,omitempty"`
}

func (o ExportRoutePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportRoutePolicy struct{}"
	}

	return strings.Join([]string{"ExportRoutePolicy", string(data)}, " ")
}
