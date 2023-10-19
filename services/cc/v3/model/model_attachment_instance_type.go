package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentInstanceType 连接类型。
type AttachmentInstanceType struct {
	AttachmentInstanceType *AttachmentInstanceTypeEnum `json:"attachment_instance_type"`
}

func (o AttachmentInstanceType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentInstanceType struct{}"
	}

	return strings.Join([]string{"AttachmentInstanceType", string(data)}, " ")
}
