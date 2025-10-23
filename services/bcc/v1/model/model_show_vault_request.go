package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVaultRequest Request Object
type ShowVaultRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// VaultId
	VaultId string `json:"vault_id"`
}

func (o ShowVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultRequest", string(data)}, " ")
}
