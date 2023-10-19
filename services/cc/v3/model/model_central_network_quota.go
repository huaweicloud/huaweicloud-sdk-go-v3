package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkQuota 租户配额详情
type CentralNetworkQuota struct {
	QuotaKey *CentralNetworkQuotaKeyEnum `json:"quota_key"`

	// 配额大小。
	QuotaLimit int32 `json:"quota_limit"`

	// 已使用配额。
	Used int32 `json:"used"`

	// 配额值的单位。
	Unit string `json:"unit"`
}

func (o CentralNetworkQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkQuota struct{}"
	}

	return strings.Join([]string{"CentralNetworkQuota", string(data)}, " ")
}
