package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTransferRequest struct {
	Body *UpdateTransferRequestBody `json:"body,omitempty"`
}

func (o UpdateTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransferRequest struct{}"
	}

	return strings.Join([]string{"UpdateTransferRequest", string(data)}, " ")
}
