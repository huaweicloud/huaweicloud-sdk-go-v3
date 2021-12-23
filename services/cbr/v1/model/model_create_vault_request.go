package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVaultRequest struct {
	Body *VaultCreateReq `json:"body,omitempty"`
}

func (o CreateVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultRequest struct{}"
	}

	return strings.Join([]string{"CreateVaultRequest", string(data)}, " ")
}
