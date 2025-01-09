package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopSubResourcesRequest Request Object
type DeleteDesktopSubResourcesRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *DeleteDesktopSubResourcesReq `json:"body,omitempty"`
}

func (o DeleteDesktopSubResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopSubResourcesRequest struct{}"
	}

	return strings.Join([]string{"DeleteDesktopSubResourcesRequest", string(data)}, " ")
}
