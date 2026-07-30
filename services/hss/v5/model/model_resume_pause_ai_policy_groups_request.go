package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumePauseAiPolicyGroupsRequest Request Object
type ResumePauseAiPolicyGroupsRequest struct {
	Body *ResumePauseAiPolicyGroupsRequestInfo `json:"body,omitempty"`
}

func (o ResumePauseAiPolicyGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePauseAiPolicyGroupsRequest struct{}"
	}

	return strings.Join([]string{"ResumePauseAiPolicyGroupsRequest", string(data)}, " ")
}
