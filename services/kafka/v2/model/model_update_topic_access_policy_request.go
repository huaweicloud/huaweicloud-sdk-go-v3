package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateTopicAccessPolicyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *UpdateTopicAccessPolicyReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateTopicAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicAccessPolicyRequest", string(data)}, " ")
}
