package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 存储库批量修改参数体
type BatchUpdateVaultRequestBody struct {
	Vault *VaultBatchUpdate `json:"vault"`
}

func (o BatchUpdateVaultRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateVaultRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateVaultRequestBody", string(data)}, " ")
}
