package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateResourceSharePermissionRequest struct {

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	Body *AssociatePermissionReqBody `json:"body,omitempty"`
}

func (o AssociateResourceSharePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResourceSharePermissionRequest struct{}"
	}

	return strings.Join([]string{"AssociateResourceSharePermissionRequest", string(data)}, " ")
}
