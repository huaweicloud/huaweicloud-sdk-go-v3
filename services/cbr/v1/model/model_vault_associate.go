package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultAssociate struct {

	// 目标vault ID , 只有设置复制策略时使用，而且必传
	DestinationVaultId *string `json:"destination_vault_id,omitempty" xml:"destination_vault_id"`

	// 策略ID。policy_id字段与add_policy_ids字段在一次请求中有且只有一个。
	PolicyId *string `json:"policy_id,omitempty" xml:"policy_id"`

	// 多策略场景下，绑定新策略的id列表。policy_id字段与add_policy_ids字段在一次请求中有且只有一个。
	AddPolicyIds *[]string `json:"add_policy_ids,omitempty" xml:"add_policy_ids"`
}

func (o VaultAssociate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultAssociate struct{}"
	}

	return strings.Join([]string{"VaultAssociate", string(data)}, " ")
}
