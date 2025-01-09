package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopRequest Request Object
type ResizeDesktopRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *ResizeDesktopReq `json:"body,omitempty"`
}

func (o ResizeDesktopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopRequest struct{}"
	}

	return strings.Join([]string{"ResizeDesktopRequest", string(data)}, " ")
}
