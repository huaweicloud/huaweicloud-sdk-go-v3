package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppServersRequest Request Object
type CreateAppServersRequest struct {

	// 交易组件调用时下发的关联ID。
	XLinkedId *string `json:"X-Linked-Id,omitempty"`

	// CBC接口回调时，请求头里带上的业务ID 包周期场景必填 按需场景无。
	ServiceTransactionId *string `json:"Service-Transaction-Id,omitempty"`

	Body *CreateAppServerReq `json:"body,omitempty"`
}

func (o CreateAppServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppServersRequest struct{}"
	}

	return strings.Join([]string{"CreateAppServersRequest", string(data)}, " ")
}
