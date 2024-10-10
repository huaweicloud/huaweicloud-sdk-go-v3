package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WafGeoIpRule struct {

	// 地理位置
	Geoip *string `json:"geoip,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 防护区域，0-大陆，1-海外
	OverseasType *int32 `json:"overseas_type,omitempty"`

	// 添加时间
	Timestamp *int32 `json:"timestamp,omitempty"`

	// 防护动作 0-阻断，1-放行，2-仅记录
	White *int32 `json:"white,omitempty"`
}

func (o WafGeoIpRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafGeoIpRule struct{}"
	}

	return strings.Join([]string{"WafGeoIpRule", string(data)}, " ")
}
