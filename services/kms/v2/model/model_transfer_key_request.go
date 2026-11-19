package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferKeyRequest Request Object
type TransferKeyRequest struct {
	Body *TransferKeyRequestBody `json:"body,omitempty"`
}

func (o TransferKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferKeyRequest struct{}"
	}

	return strings.Join([]string{"TransferKeyRequest", string(data)}, " ")
}
