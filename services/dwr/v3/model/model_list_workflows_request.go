package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowsRequest Request Object
type ListWorkflowsRequest struct {

	// 工作流的名称前缀。名称必须以字母或数字开头，只能由字母、数字、下划线和中划线组成，长度小于等于64个字符。
	Prefix *string `json:"prefix,omitempty"`

	// 查询的起始位置。start大于等于1，最大1000，不设置则取默认值1。
	Offset *int32 `json:"offset,omitempty"`

	// 请求返回的最大记录条数。limit取值最小1，最大1000，不设置则取默认值10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListWorkflowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowsRequest", string(data)}, " ")
}
