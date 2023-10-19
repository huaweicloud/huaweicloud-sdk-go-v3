package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkQuotaKey 带宽值，单位Mbps。
type CentralNetworkQuotaKey struct {
	QuotaKey *CentralNetworkQuotaKeyEnum `json:"quota_key"`
}

func (o CentralNetworkQuotaKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkQuotaKey struct{}"
	}

	return strings.Join([]string{"CentralNetworkQuotaKey", string(data)}, " ")
}
