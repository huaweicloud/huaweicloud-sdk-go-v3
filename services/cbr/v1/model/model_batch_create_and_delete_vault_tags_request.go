package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateAndDeleteVaultTagsRequest struct {

	// 资源id
	VaultId string `json:"vault_id"`

	Body *BulkCreateAndDeleteVaultTagsReq `json:"body,omitempty"`
}

func (o BatchCreateAndDeleteVaultTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateAndDeleteVaultTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateAndDeleteVaultTagsRequest", string(data)}, " ")
}
