package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteStackRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 用户希望操作的资源栈名
	StackName string `json:"stack_name"`

	// 用户希望描述的栈的Id。若stack_name和stack_id同时存在，则IaC会检查是否两个匹配，否则返回400
	StackId *string `json:"stack_id,omitempty"`

	// 执行操作者的名字，将用做未来的审计工作。
	Executor *string `json:"executor,omitempty"`
}

func (o DeleteStackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStackRequest struct{}"
	}

	return strings.Join([]string{"DeleteStackRequest", string(data)}, " ")
}
