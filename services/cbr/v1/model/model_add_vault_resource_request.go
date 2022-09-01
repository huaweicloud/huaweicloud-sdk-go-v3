package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddVaultResourceRequest struct {

	// 存储库ID
	VaultId string `json:"vault_id" xml:"vault_id"`

	Body *VaultAddResourceReq `json:"body,omitempty" xml:"body"`
}

func (o AddVaultResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVaultResourceRequest struct{}"
	}

	return strings.Join([]string{"AddVaultResourceRequest", string(data)}, " ")
}
