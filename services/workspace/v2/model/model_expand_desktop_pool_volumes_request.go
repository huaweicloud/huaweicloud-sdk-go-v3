package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopPoolVolumesRequest Request Object
type ExpandDesktopPoolVolumesRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *ExpandDesktopPoolVolumesReq `json:"body,omitempty"`
}

func (o ExpandDesktopPoolVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopPoolVolumesRequest struct{}"
	}

	return strings.Join([]string{"ExpandDesktopPoolVolumesRequest", string(data)}, " ")
}
