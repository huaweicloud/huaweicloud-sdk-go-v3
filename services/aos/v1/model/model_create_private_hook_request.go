package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateHookRequest Request Object
type CreatePrivateHookRequest struct {

	// 用户指定的，对于此请求的唯一Id，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	Body *CreatePrivateHookRequestBody `json:"body,omitempty"`
}

func (o CreatePrivateHookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateHookRequest struct{}"
	}

	return strings.Join([]string{"CreatePrivateHookRequest", string(data)}, " ")
}
