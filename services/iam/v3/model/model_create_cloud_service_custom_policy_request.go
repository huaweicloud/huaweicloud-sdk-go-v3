package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudServiceCustomPolicyRequest Request Object
type CreateCloudServiceCustomPolicyRequest struct {
	Body *CreateCloudServiceCustomPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateCloudServiceCustomPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudServiceCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudServiceCustomPolicyRequest", string(data)}, " ")
}
