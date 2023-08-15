package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcAttachmentRequest Request Object
type DeleteVpcAttachmentRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// VPC连接ID
	VpcAttachmentId string `json:"vpc_attachment_id"`
}

func (o DeleteVpcAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcAttachmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpcAttachmentRequest", string(data)}, " ")
}
