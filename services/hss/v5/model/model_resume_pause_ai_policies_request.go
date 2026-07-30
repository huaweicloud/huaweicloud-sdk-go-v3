package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumePauseAiPoliciesRequest Request Object
type ResumePauseAiPoliciesRequest struct {
	Body *ResumePauseAiPoliciesRequestInfo `json:"body,omitempty"`
}

func (o ResumePauseAiPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePauseAiPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ResumePauseAiPoliciesRequest", string(data)}, " ")
}
