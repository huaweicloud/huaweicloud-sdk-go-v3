package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumePauseAiPoliciesResponse Response Object
type ResumePauseAiPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResumePauseAiPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePauseAiPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ResumePauseAiPoliciesResponse", string(data)}, " ")
}
