package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcAttachmentDetail struct {

	// 东西向防护添加的防护vpc的id
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o VpcAttachmentDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcAttachmentDetail struct{}"
	}

	return strings.Join([]string{"VpcAttachmentDetail", string(data)}, " ")
}
