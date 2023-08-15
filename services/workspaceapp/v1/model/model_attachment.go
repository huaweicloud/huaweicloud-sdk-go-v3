package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Attachment 配置关联成员
type Attachment struct {

	// 目标
	Attach string `json:"attach"`

	AttachType *AttachType `json:"attach_type"`
}

func (o Attachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attachment struct{}"
	}

	return strings.Join([]string{"Attachment", string(data)}, " ")
}
