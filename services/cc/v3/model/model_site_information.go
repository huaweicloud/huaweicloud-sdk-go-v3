package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteInformation 分支信息。
type SiteInformation struct {

	// RegionID。
	RegionId string `json:"region_id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	GatewayType *GatewayTypeEnum `json:"gateway_type"`

	// 实例ID。
	GatewayId string `json:"gateway_id"`

	// 站点编码定义
	SiteCode string `json:"site_code"`

	// 网络实例BGP协议的AS号。
	Asn int64 `json:"asn"`
}

func (o SiteInformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteInformation struct{}"
	}

	return strings.Join([]string{"SiteInformation", string(data)}, " ")
}
