package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransferRequest Request Object
type CreateTransferRequest struct {
	Body *CreateTransferRequestBody `json:"body,omitempty"`
}

func (o CreateTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferRequest struct{}"
	}

	return strings.Join([]string{"CreateTransferRequest", string(data)}, " ")
}
