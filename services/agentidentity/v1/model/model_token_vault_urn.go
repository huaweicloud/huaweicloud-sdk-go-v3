package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TokenVaultUrn TokenVault 对象统一资源标识（URN）。
type TokenVaultUrn struct {
}

func (o TokenVaultUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenVaultUrn struct{}"
	}

	return strings.Join([]string{"TokenVaultUrn", string(data)}, " ")
}
