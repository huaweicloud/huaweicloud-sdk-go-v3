package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowInstanceResponse Response Object
type ListWorkflowInstanceResponse struct {

	// 满足条件的运行实例个数。
	Count *int32 `json:"count,omitempty"`

	// 实例信息列表
	Executions *[]Execution `json:"executions,omitempty"`

	// 表明是否本次返回的结果列表被截断。true：表示本次没有返回全部结果。false：表示本次已经返回了全部结果。
	IsTruncated *bool `json:"is_truncated,omitempty"`

	// 如果本次没有返回全部结果，响应请求中将包含此字段，用于标明本次请求列举到的最后一个工作流实例。后续请求可以指定Marker等于该值来列举剩余的工作流实例。如果is_truncated为false，该字段不会返回。
	NextOffset *int32 `json:"next_offset,omitempty"`

	XRequestId *string `json:"X-Request-Id,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date *string `json:"Date,omitempty"`

	ContentType    *string `json:"Content-Type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListWorkflowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowInstanceResponse", string(data)}, " ")
}
