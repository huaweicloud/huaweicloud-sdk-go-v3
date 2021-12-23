package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteVaultRequest struct {
	// 存储库ID

	VaultId string `json:"vault_id"`
}

func (o DeleteVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVaultRequest struct{}"
	}

	return strings.Join([]string{"DeleteVaultRequest", string(data)}, " ")
}
