package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetBackupPolicyRequest Request Object
type SetBackupPolicyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SetBackupPolicyRequestBody `json:"body,omitempty"`
}

func (o SetBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetBackupPolicyRequest", string(data)}, " ")
}
