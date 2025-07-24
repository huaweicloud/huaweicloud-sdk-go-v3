package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateObsTargetAttributesRequest Request Object
type UpdateObsTargetAttributesRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	// 绑定关系ID
	TargetId string `json:"target_id"`

	Body *UpdateObsTargetAttributesRequestBody `json:"body,omitempty"`
}

func (o UpdateObsTargetAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateObsTargetAttributesRequest struct{}"
	}

	return strings.Join([]string{"UpdateObsTargetAttributesRequest", string(data)}, " ")
}
