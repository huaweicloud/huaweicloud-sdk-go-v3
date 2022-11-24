package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVaultResourceInstancesRequest struct {
	Body *VaultResourceInstancesReq `json:"body,omitempty"`
}

func (o ShowVaultResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVaultResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowVaultResourceInstancesRequest", string(data)}, " ")
}
