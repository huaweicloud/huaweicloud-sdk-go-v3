package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExportPrivateKeyRequest Request Object
type BatchExportPrivateKeyRequest struct {
	Body *[]Keypairs `json:"body,omitempty"`
}

func (o BatchExportPrivateKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportPrivateKeyRequest struct{}"
	}

	return strings.Join([]string{"BatchExportPrivateKeyRequest", string(data)}, " ")
}
