package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddVolumesRequest Request Object
type AddVolumesRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *AddDesktopsVolumesReq `json:"body,omitempty"`
}

func (o AddVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVolumesRequest struct{}"
	}

	return strings.Join([]string{"AddVolumesRequest", string(data)}, " ")
}
