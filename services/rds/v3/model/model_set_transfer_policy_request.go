package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTransferPolicyRequest Request Object
type SetTransferPolicyRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	Body *SetTransferPolicyRequestBody `json:"body,omitempty"`
}

func (o SetTransferPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTransferPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetTransferPolicyRequest", string(data)}, " ")
}
