package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchLogoffDesktopsRequest Request Object
type BatchLogoffDesktopsRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *LogoffDesktopsReq `json:"body,omitempty"`
}

func (o BatchLogoffDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchLogoffDesktopsRequest struct{}"
	}

	return strings.Join([]string{"BatchLogoffDesktopsRequest", string(data)}, " ")
}
