package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpcAttachmentResponse Response Object
type DeleteVpcAttachmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpcAttachmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcAttachmentResponse", string(data)}, " ")
}
