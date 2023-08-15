package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportRoutePolicy 入方向路由策略
type ImportRoutePolicy struct {

	// 入方向Ipv4协议路由策略id
	ImportPolicyId *string `json:"import_policy_id,omitempty"`
}

func (o ImportRoutePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportRoutePolicy struct{}"
	}

	return strings.Join([]string{"ImportRoutePolicy", string(data)}, " ")
}
