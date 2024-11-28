package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopPoolRequest Request Object
type ExpandDesktopPoolRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *ExpandDesktopPoolReq `json:"body,omitempty"`
}

func (o ExpandDesktopPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopPoolRequest struct{}"
	}

	return strings.Join([]string{"ExpandDesktopPoolRequest", string(data)}, " ")
}
