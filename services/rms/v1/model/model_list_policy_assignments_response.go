package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPolicyAssignmentsResponse struct {

	// 规则列表
	Value *[]PolicyAssignment `json:"value,omitempty" xml:"value"`

	PageInfo       *PageInfo `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPolicyAssignmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyAssignmentsResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyAssignmentsResponse", string(data)}, " ")
}
