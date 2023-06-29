package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateVaultRequest Request Object
type BatchUpdateVaultRequest struct {
	Body *BatchUpdateVaultRequestBody `json:"body,omitempty"`
}

func (o BatchUpdateVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateVaultRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateVaultRequest", string(data)}, " ")
}
