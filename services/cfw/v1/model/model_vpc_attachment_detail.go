package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcAttachmentDetail struct {

	// **参数解释**： 东西向防护添加的防护VPC的ID **取值范围**： 不涉及
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o VpcAttachmentDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcAttachmentDetail struct{}"
	}

	return strings.Join([]string{"VpcAttachmentDetail", string(data)}, " ")
}
