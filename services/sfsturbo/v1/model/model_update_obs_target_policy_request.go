package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObsTargetPolicyRequest Request Object
type UpdateObsTargetPolicyRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	// 绑定关系ID
	TargetId string `json:"target_id"`

	Body *UpdateObsTargetPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateObsTargetPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObsTargetPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateObsTargetPolicyRequest", string(data)}, " ")
}
