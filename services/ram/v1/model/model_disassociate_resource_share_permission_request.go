package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateResourceSharePermissionRequest struct {

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	Body *DisassociatePermissionReqBody `json:"body,omitempty"`
}

func (o DisassociateResourceSharePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResourceSharePermissionRequest struct{}"
	}

	return strings.Join([]string{"DisassociateResourceSharePermissionRequest", string(data)}, " ")
}
