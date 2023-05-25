package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePostPaidVaultRequest struct {
	Body *VaultOrderCreateReqs `json:"body,omitempty"`
}

func (o CreatePostPaidVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidVaultRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPaidVaultRequest", string(data)}, " ")
}
