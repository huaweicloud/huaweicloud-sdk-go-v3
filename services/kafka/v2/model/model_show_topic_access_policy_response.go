package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopicAccessPolicyResponse Response Object
type ShowTopicAccessPolicyResponse struct {

	// Topic名称。
	Name *string `json:"name,omitempty"`

	// Topic类型。
	TopicType *int32 `json:"topic_type,omitempty"`

	// 权限列表。
	Policies       *[]PolicyEntity `json:"policies,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowTopicAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowTopicAccessPolicyResponse", string(data)}, " ")
}
