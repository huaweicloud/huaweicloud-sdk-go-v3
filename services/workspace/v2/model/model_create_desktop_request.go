package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopRequest Request Object
type CreateDesktopRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *CreateDesktopReq `json:"body,omitempty"`
}

func (o CreateDesktopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopRequest struct{}"
	}

	return strings.Join([]string{"CreateDesktopRequest", string(data)}, " ")
}
