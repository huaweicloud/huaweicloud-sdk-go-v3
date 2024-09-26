package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CentralNetworkErInstance struct {

	// 实例ID。
	Id string `json:"id"`

	// 企业路由器的ID。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 网络实例BGP协议的AS号。
	Asn int64 `json:"asn"`

	// 站点编码定义
	SiteCode string `json:"site_code"`
}

func (o CentralNetworkErInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkErInstance struct{}"
	}

	return strings.Join([]string{"CentralNetworkErInstance", string(data)}, " ")
}
