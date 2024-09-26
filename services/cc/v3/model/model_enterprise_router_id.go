package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseRouterId 企业路由器的ID。
type EnterpriseRouterId struct {

	// 企业路由器的ID。
	EnterpriseRouterId string `json:"enterprise_router_id"`
}

func (o EnterpriseRouterId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseRouterId struct{}"
	}

	return strings.Join([]string{"EnterpriseRouterId", string(data)}, " ")
}
