package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ContinueRollbackStackRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 用户希望操作的资源栈名称
	StackName string `json:"stack_name"`

	Body *ContinueRollbackStackRequestBody `json:"body,omitempty"`
}

func (o ContinueRollbackStackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContinueRollbackStackRequest struct{}"
	}

	return strings.Join([]string{"ContinueRollbackStackRequest", string(data)}, " ")
}
