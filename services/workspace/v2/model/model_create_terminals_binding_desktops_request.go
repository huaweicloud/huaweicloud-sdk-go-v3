package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTerminalsBindingDesktopsRequest Request Object
type CreateTerminalsBindingDesktopsRequest struct {
	Body *CreateTerminalsBindingDesktopsRequestBody `json:"body,omitempty"`
}

func (o CreateTerminalsBindingDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTerminalsBindingDesktopsRequest struct{}"
	}

	return strings.Join([]string{"CreateTerminalsBindingDesktopsRequest", string(data)}, " ")
}
