package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddVaultResourceRequest Request Object
type AddVaultResourceRequest struct {

	// 存储库ID
	VaultId string `json:"vault_id"`

	Body *VaultAddResourceReq `json:"body,omitempty"`
}

func (o AddVaultResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"AddVaultResourceRequest", string(data)}, " ")
}
