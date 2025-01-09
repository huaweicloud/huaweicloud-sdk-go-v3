package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandVolumesRequest Request Object
type ExpandVolumesRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *ExpandDesktopsVolumesReq `json:"body,omitempty"`
}

func (o ExpandVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandVolumesRequest struct{}"
	}

	return strings.Join([]string{"ExpandVolumesRequest", string(data)}, " ")
}
