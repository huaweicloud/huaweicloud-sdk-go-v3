package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectInfo 防护小结信息
type ProtectInfo struct {
	CoverAreaInfo *ProtectInfoCoverAreaInfo `json:"cover_area_info,omitempty"`

	ConfigInfo *ProtectInfoConfigInfo `json:"config_info,omitempty"`

	QuotaInfo *ProtectInfoQuotaInfo `json:"quota_info,omitempty"`
}

func (o ProtectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectInfo struct{}"
	}

	return strings.Join([]string{"ProtectInfo", string(data)}, " ")
}
