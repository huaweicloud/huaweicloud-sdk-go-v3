package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyJobsRequest Request Object
type ListPolicyJobsRequest struct {

	// 指定的任务类型，没有则使用默认值EnablePolicy
	Kind *string `json:"kind,omitempty"`
}

func (o ListPolicyJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyJobsRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyJobsRequest", string(data)}, " ")
}
