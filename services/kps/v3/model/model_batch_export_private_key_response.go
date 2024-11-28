package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExportPrivateKeyResponse Response Object
type BatchExportPrivateKeyResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchExportPrivateKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExportPrivateKeyResponse struct{}"
	}

	return strings.Join([]string{"BatchExportPrivateKeyResponse", string(data)}, " ")
}
