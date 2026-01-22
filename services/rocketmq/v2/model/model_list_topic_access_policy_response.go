package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopicAccessPolicyResponse Response Object
type ListTopicAccessPolicyResponse struct {

	// **参数解释**： 用户列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Policies *[]ListAccessPolicyRespPolicies `json:"policies,omitempty"`

	// **参数解释**： 总用户个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total float32 `json:"total,omitempty"`

	// **参数解释**： 主题或消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTopicAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListTopicAccessPolicyResponse", string(data)}, " ")
}
