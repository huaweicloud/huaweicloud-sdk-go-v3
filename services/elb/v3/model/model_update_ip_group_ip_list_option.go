package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新IP地址组IP列表请求参数。
type UpdateIpGroupIpListOption struct {
	// IP地址组的名称

	Name *string `json:"name,omitempty"`
	// IP地址组中包含的IP列表。

	IpList *[]UpadateIpGroupIpOption `json:"ip_list,omitempty"`
	// IP地址组的描述信息

	Description *string `json:"description,omitempty"`
}

func (o UpdateIpGroupIpListOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupIpListOption struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupIpListOption", string(data)}, " ")
}
