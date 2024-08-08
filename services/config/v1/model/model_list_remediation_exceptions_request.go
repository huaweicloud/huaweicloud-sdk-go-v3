package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemediationExceptionsRequest Request Object
type ListRemediationExceptionsRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o ListRemediationExceptionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemediationExceptionsRequest struct{}"
	}

	return strings.Join([]string{"ListRemediationExceptionsRequest", string(data)}, " ")
}
