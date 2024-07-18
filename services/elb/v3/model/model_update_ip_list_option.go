package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpListOption 更新IP地址组IP列表请求参数。
type UpdateIpListOption struct {

	// 参数解释：IP地址组的名称
	Name *string `json:"name,omitempty"`

	// 参数解释：IP地址组中包含的IP列表。
	IpList *[]UpadateIpGroupIpOption `json:"ip_list,omitempty"`

	// 参数解释：IP地址组的描述信息
	Description *string `json:"description,omitempty"`
}

func (o UpdateIpListOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListOption struct{}"
	}

	return strings.Join([]string{"UpdateIpListOption", string(data)}, " ")
}
