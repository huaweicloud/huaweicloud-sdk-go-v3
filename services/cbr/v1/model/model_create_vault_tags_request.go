package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVaultTagsRequest Request Object
type CreateVaultTagsRequest struct {

	// 资源id
	VaultId string `json:"vault_id"`

	Body *VaultTagsCreateReq `json:"body,omitempty"`
}

func (o CreateVaultTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateVaultTagsRequest", string(data)}, " ")
}
