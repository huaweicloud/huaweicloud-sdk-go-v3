package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateVpcAttachmentRequest struct {

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	Body *CreateVpcAttachmentBody `json:"body,omitempty"`
}

func (o CreateVpcAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcAttachmentRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcAttachmentRequest", string(data)}, " ")
}
