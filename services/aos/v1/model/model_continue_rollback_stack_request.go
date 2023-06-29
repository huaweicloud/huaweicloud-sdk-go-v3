package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContinueRollbackStackRequest Request Object
type ContinueRollbackStackRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 资源栈的名称。此名字在domain_id+区域+project_id下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
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
