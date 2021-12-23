package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVaultRequest struct {
	// 存储库ID

	VaultId string `json:"vault_id"`
}

func (o ShowVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultRequest", string(data)}, " ")
}
