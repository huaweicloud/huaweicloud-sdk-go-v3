package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpInfo IP地址组中的IP网段。
type IpInfo struct {

	// IP地址组中的IP网段，cidr格式。
	Cidr *string `json:"cidr,omitempty"`

	// IP地址组中的IP网段描述。
	Description *string `json:"description,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o IpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpInfo struct{}"
	}

	return strings.Join([]string{"IpInfo", string(data)}, " ")
}
