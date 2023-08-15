package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseRouterAz struct {

	// 企业路由器实例所在的可用区
	AvailabilityZoneIds []string `json:"availability_zone_ids"`
}

func (o EnterpriseRouterAz) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseRouterAz struct{}"
	}

	return strings.Join([]string{"EnterpriseRouterAz", string(data)}, " ")
}
