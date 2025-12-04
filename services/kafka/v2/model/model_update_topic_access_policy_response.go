package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopicAccessPolicyResponse Response Object
type UpdateTopicAccessPolicyResponse struct {

	// **参数解释**： 后台任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTopicAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyResponse", string(data)}, " ")
}
