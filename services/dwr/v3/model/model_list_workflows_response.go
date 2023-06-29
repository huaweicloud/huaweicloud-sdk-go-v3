package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowsResponse Response Object
type ListWorkflowsResponse struct {

	// 列表总条数。
	Count *int32 `json:"count,omitempty"`

	// 工作流模板列表信息。
	Graphs *[]GraphItem `json:"graphs,omitempty"`

	// 下一次查询的起始位置。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// 表明是否本次返回的ListWorkflow结果列表被截断。“true”表示本次没有返回全部结果；“false”表示本次已经返回了全部结果。
	IsTruncated *bool `json:"is_truncated,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListWorkflowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowsResponse", string(data)}, " ")
}
