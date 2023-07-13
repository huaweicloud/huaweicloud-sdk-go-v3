package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpcAttachmentRequest Request Object
type ShowVpcAttachmentRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// VPC连接ID
	VpcAttachmentId string `json:"vpc_attachment_id"`
}

func (o ShowVpcAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpcAttachmentRequest struct{}"
	}

	return strings.Join([]string{"ShowVpcAttachmentRequest", string(data)}, " ")
}
