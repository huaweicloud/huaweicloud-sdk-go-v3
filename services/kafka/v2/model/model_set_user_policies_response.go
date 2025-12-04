package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetUserPoliciesResponse Response Object
type SetUserPoliciesResponse struct {

	// **参数解释**： 后台任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetUserPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetUserPoliciesResponse struct{}"
	}

	return strings.Join([]string{"SetUserPoliciesResponse", string(data)}, " ")
}
