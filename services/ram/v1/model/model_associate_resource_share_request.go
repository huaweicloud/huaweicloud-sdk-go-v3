package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateResourceShareRequest struct {

	// 资源共享实例的ID。
	ResourceShareId string `json:"resource_share_id"`

	Body *ResourceShareAssociationReqBody `json:"body,omitempty"`
}

func (o AssociateResourceShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResourceShareRequest struct{}"
	}

	return strings.Join([]string{"AssociateResourceShareRequest", string(data)}, " ")
}
