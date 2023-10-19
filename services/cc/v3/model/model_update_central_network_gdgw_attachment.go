package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkGdgwAttachment 更新中心网络GDGW附件的属性详情。
type UpdateCentralNetworkGdgwAttachment struct {

	// 实例名字。
	Name *string `json:"name,omitempty"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`
}

func (o UpdateCentralNetworkGdgwAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkGdgwAttachment struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkGdgwAttachment", string(data)}, " ")
}
