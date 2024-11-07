package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DirectedEdge struct {

	// 实例ID。
	Id string `json:"id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 实例ID。
	GatewayId string `json:"gateway_id"`

	GatewayType *GatewayTypeEnum `json:"gateway_type"`

	// 站点编码定义
	SiteCode string `json:"site_code"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`
}

func (o DirectedEdge) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectedEdge struct{}"
	}

	return strings.Join([]string{"DirectedEdge", string(data)}, " ")
}
