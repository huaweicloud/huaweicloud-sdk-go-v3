package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetVaultResourceRequest Request Object
type SetVaultResourceRequest struct {

	// 存储库id
	VaultId string `json:"vault_id"`

	Body *VaultSetResourceReq `json:"body,omitempty"`
}

func (o SetVaultResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"SetVaultResourceRequest", string(data)}, " ")
}
