package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TokenVaultId The unique identifier of the token vault.
type TokenVaultId struct {
}

func (o TokenVaultId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenVaultId struct{}"
	}

	return strings.Join([]string{"TokenVaultId", string(data)}, " ")
}
