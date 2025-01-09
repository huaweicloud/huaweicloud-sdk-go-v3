package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyDesktopsInternetRequest Request Object
type ApplyDesktopsInternetRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *ApplyDesktopsInternetReq `json:"body,omitempty"`
}

func (o ApplyDesktopsInternetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyDesktopsInternetRequest struct{}"
	}

	return strings.Join([]string{"ApplyDesktopsInternetRequest", string(data)}, " ")
}
