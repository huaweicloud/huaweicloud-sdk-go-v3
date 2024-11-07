package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteNetworkQuotaKey 带宽值，单位Mbps。
type SiteNetworkQuotaKey struct {
	QuotaKey *SiteNetworkQuotaKeyEnum `json:"quota_key"`
}

func (o SiteNetworkQuotaKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteNetworkQuotaKey struct{}"
	}

	return strings.Join([]string{"SiteNetworkQuotaKey", string(data)}, " ")
}
