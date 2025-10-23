package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigureVaultItem 配置保护的存储库实体
type ConfigureVaultItem struct {

	// CBR存储库ID
	VaultId string `json:"vault_id"`

	// 策略ID
	PolicyId string `json:"policy_id"`

	// 是否加密
	Encrypted *bool `json:"encrypted,omitempty"`

	// 是否备份锁定
	Worm *bool `json:"worm,omitempty"`

	// 策略是否启用
	PolicyEnable bool `json:"policy_enable"`
}

func (o ConfigureVaultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigureVaultItem struct{}"
	}

	return strings.Join([]string{"ConfigureVaultItem", string(data)}, " ")
}
