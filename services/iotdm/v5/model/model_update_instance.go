package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstance struct {

	// **参数说明**：实例名称 **取值范围**：由中文字符，英文字母、数字及“_”、“-”组成，且长度为[1-64]个字符。
	Name *string `json:"name,omitempty"`

	// **参数说明**：设备接入实例的描述信息。 **取值范围**：由中文，字母，数字，句号，逗号，下划线（“_”），中划线（“-”），空格组成，且长度为[1-256]个字符。
	Description *string `json:"description,omitempty"`

	OperateWindow *OperateWindow `json:"operate_window,omitempty"`

	ForwardingInfo *UpdateForwardingInfo `json:"forwarding_info,omitempty"`

	AccessInfo *UpdateAccessInfo `json:"access_info,omitempty"`
}

func (o UpdateInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstance struct{}"
	}

	return strings.Join([]string{"UpdateInstance", string(data)}, " ")
}
