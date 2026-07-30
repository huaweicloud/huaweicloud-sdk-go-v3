package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumePauseAiPolicyGroupsResponse Response Object
type ResumePauseAiPolicyGroupsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResumePauseAiPolicyGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePauseAiPolicyGroupsResponse struct{}"
	}

	return strings.Join([]string{"ResumePauseAiPolicyGroupsResponse", string(data)}, " ")
}
