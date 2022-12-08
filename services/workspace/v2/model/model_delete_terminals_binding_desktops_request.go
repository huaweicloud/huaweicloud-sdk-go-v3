package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteTerminalsBindingDesktopsRequest struct {
	Body *DeleteTerminalsBindingDesktopsRequestBody `json:"body,omitempty"`
}

func (o DeleteTerminalsBindingDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTerminalsBindingDesktopsRequest struct{}"
	}

	return strings.Join([]string{"DeleteTerminalsBindingDesktopsRequest", string(data)}, " ")
}
