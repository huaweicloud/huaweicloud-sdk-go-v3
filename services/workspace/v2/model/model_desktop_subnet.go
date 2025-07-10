package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopSubnet 子网信息。
type DesktopSubnet struct {

	// 桌面所在子网Id。
	Id *string `json:"id,omitempty"`

	// 桌面所在子网名称。
	Name *string `json:"name,omitempty"`

	// 桌面所在子网网段。
	Cidr *string `json:"cidr,omitempty"`
}

func (o DesktopSubnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopSubnet struct{}"
	}

	return strings.Join([]string{"DesktopSubnet", string(data)}, " ")
}
