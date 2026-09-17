package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteInfo 站点信息
type SiteInfo struct {

	// 运营站点编码
	SiteCode *string `json:"site_code,omitempty"`

	// 运营站点名称
	SiteName *string `json:"site_name,omitempty"`

	// 云服务区信息列表
	Regions *[]RegionInfo `json:"regions,omitempty"`
}

func (o SiteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteInfo struct{}"
	}

	return strings.Join([]string{"SiteInfo", string(data)}, " ")
}
