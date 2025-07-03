package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateNetworkInterfaceOption 网络接口，网卡
type TemplateNetworkInterfaceOption struct {

	// 子网ID
	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	Attachment *TemplateNetworkInterfaceAttachmentOption `json:"attachment,omitempty"`
}

func (o TemplateNetworkInterfaceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateNetworkInterfaceOption struct{}"
	}

	return strings.Join([]string{"TemplateNetworkInterfaceOption", string(data)}, " ")
}
