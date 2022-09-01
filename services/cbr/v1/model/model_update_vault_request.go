package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateVaultRequest struct {

	// 存储库ID
	VaultId string `json:"vault_id" xml:"vault_id"`

	Body *VaultUpdateReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVaultRequest struct{}"
	}

	return strings.Join([]string{"UpdateVaultRequest", string(data)}, " ")
}
