package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceShareRequest Request Object
type UpdateResourceShareRequest struct {

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	Body *UpdateResourceShareReqBody `json:"body,omitempty"`
}

func (o UpdateResourceShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceShareRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceShareRequest", string(data)}, " ")
}
