package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListConsumeGroupAccessPolicyResponse struct {

	// 用户列表。
	Policies *[]ListAccessPolicyRespPolicies `json:"policies,omitempty" xml:"policies"`

	// 总用户个数。
	Total float32 `json:"total,omitempty" xml:"total"`

	// 主题或消费组名称。
	Name           *string `json:"name,omitempty" xml:"name"`
	HttpStatusCode int     `json:"-"`
}

func (o ListConsumeGroupAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumeGroupAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListConsumeGroupAccessPolicyResponse", string(data)}, " ")
}
