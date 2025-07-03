package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateNetworkInterfaceAttachmentOption 网络接口，网卡
type TemplateNetworkInterfaceAttachmentOption struct {

	// 加载顺序, 0代表主网卡
	DeviceIndex *int32 `json:"device_index,omitempty"`
}

func (o TemplateNetworkInterfaceAttachmentOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateNetworkInterfaceAttachmentOption struct{}"
	}

	return strings.Join([]string{"TemplateNetworkInterfaceAttachmentOption", string(data)}, " ")
}
