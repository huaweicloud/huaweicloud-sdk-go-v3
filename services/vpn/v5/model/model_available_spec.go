package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableSpec struct {

	// VPN网关规格
	Flavor *string `json:"flavor,omitempty"`

	// 关联模式
	AttachmentType *string `json:"attachment_type,omitempty"`

	// 网关的IP协议版本
	IpVersion *string `json:"ip_version,omitempty"`
}

func (o AvailableSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableSpec struct{}"
	}

	return strings.Join([]string{"AvailableSpec", string(data)}, " ")
}
