package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteNetworkQuota 分支网络租户配额详情。
type SiteNetworkQuota struct {
	QuotaKey *SiteNetworkQuotaKeyEnum `json:"quota_key"`

	// 配额大小。
	QuotaLimit int32 `json:"quota_limit"`

	// 已使用配额。
	Used int32 `json:"used"`

	// 配额值的单位。
	Unit string `json:"unit"`
}

func (o SiteNetworkQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteNetworkQuota struct{}"
	}

	return strings.Join([]string{"SiteNetworkQuota", string(data)}, " ")
}
