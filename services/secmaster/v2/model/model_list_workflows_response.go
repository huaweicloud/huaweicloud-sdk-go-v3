package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowsResponse Response Object
type ListWorkflowsResponse struct {

	// 返回码
	Code *string `json:"code,omitempty"`

	// 数据总条数
	Total *int32 `json:"total,omitempty"`

	// 当前页大小
	Offset *int32 `json:"offset,omitempty"`

	// 当前页码
	Limit *int32 `json:"limit,omitempty"`

	// 请求ID
	Message *string `json:"message,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 流程信息列表
	Data *[]AopWorkflowInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListWorkflowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowsResponse", string(data)}, " ")
}
