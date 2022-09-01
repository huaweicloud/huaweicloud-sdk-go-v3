package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPolicyStatesByDomainIdRequest struct {

	// 合规状态
	ComplianceState *string `json:"compliance_state,omitempty" xml:"compliance_state"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty" xml:"resource_name"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty" xml:"marker"`
}

func (o ListPolicyStatesByDomainIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyStatesByDomainIdRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyStatesByDomainIdRequest", string(data)}, " ")
}
