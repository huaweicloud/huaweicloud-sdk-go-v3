package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteNetworkId 分支网络ID。
type SiteNetworkId struct {

	// 实例ID。
	SiteNetworkId string `json:"site_network_id"`
}

func (o SiteNetworkId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteNetworkId struct{}"
	}

	return strings.Join([]string{"SiteNetworkId", string(data)}, " ")
}
