package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 绑定策略返回体
type VaultPolicyResp struct {
	// 目标region的vault ID，仅设置复制策略时有。

	DestinationVaultId *string `json:"destination_vault_id,omitempty"`
	// 设置的策略ID

	PolicyId string `json:"policy_id"`
	// 设置策略的vault ID

	VaultId string `json:"vault_id"`
}

func (o VaultPolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultPolicyResp struct{}"
	}

	return strings.Join([]string{"VaultPolicyResp", string(data)}, " ")
}
