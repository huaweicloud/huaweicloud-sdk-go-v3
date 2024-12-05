package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PeerSiteExternal 对等体对象
type PeerSiteExternal struct {

	// 类型
	Type *string `json:"type,omitempty"`

	// 网关ID
	GatewayId *string `json:"gateway_id,omitempty"`

	// 连接ID
	LinkId *string `json:"link_id,omitempty"`

	// 局点ID
	RegionId *string `json:"region_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 站点编码
	SiteCode *string `json:"site_code,omitempty"`
}

func (o PeerSiteExternal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeerSiteExternal struct{}"
	}

	return strings.Join([]string{"PeerSiteExternal", string(data)}, " ")
}
