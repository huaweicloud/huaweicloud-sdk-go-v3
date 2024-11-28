package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildDesktopPoolRequest Request Object
type RebuildDesktopPoolRequest struct {

	// CBC接口回调时，请求头里带上的业务ID
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *RebuildDesktopPoolReq `json:"body,omitempty"`
}

func (o RebuildDesktopPoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildDesktopPoolRequest struct{}"
	}

	return strings.Join([]string{"RebuildDesktopPoolRequest", string(data)}, " ")
}
