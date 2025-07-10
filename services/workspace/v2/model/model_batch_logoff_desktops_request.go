package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchLogoffDesktopsRequest Request Object
type BatchLogoffDesktopsRequest struct {
	Body *LogoffDesktopsReq `json:"body,omitempty"`
}

func (o BatchLogoffDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchLogoffDesktopsRequest struct{}"
	}

	return strings.Join([]string{"BatchLogoffDesktopsRequest", string(data)}, " ")
}
