package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopicAccessPolicyResponse Response Object
type ListTopicAccessPolicyResponse struct {

	// 用户列表。
	Policies *[]ListAccessPolicyRespPolicies `json:"policies,omitempty"`

	// 总用户个数。
	Total float32 `json:"total,omitempty"`

	// 主题或消费组名称。
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
