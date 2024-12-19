package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateExternalPeerLinkRequestBodyPeerLinkPeerSite struct {

	// 接入网关连接对端的实例(当前ER实例)ID
	GatewayId string `json:"gateway_id"`

	// 对端实例(ER实例)归属的项目ID
	ProjectId string `json:"project_id"`

	// 归属的区域ID
	RegionId string `json:"region_id"`
}

func (o CreateExternalPeerLinkRequestBodyPeerLinkPeerSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExternalPeerLinkRequestBodyPeerLinkPeerSite struct{}"
	}

	return strings.Join([]string{"CreateExternalPeerLinkRequestBodyPeerLinkPeerSite", string(data)}, " ")
}
