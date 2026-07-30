package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyAiPolicyGroupRequest Request Object
type CopyAiPolicyGroupRequest struct {
	Body *CopyAiPolicyGroupRequestInfo `json:"body,omitempty"`
}

func (o CopyAiPolicyGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyAiPolicyGroupRequest struct{}"
	}

	return strings.Join([]string{"CopyAiPolicyGroupRequest", string(data)}, " ")
}
