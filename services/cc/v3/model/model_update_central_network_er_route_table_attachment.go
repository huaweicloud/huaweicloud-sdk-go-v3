package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkErRouteTableAttachment 更新中心网络ER路由表附件的属性详情。
type UpdateCentralNetworkErRouteTableAttachment struct {

	// 实例名字。
	Name *string `json:"name,omitempty"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`
}

func (o UpdateCentralNetworkErRouteTableAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkErRouteTableAttachment struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkErRouteTableAttachment", string(data)}, " ")
}
