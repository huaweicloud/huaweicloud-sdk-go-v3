package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type VaultMigrateResourceReq struct {

	// 目标存储库
	DestinationVaultId string `json:"destination_vault_id"`

	// 待迁移的资源ID
	ResourceIds []string `json:"resource_ids"`
}

func (o VaultMigrateResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultMigrateResourceReq struct{}"
	}

	return strings.Join([]string{"VaultMigrateResourceReq", string(data)}, " ")
}
