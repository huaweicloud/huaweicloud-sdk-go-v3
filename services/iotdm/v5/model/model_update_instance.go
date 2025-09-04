package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstance struct {

	// **参数说明**：实例名称。 **取值范围**：由中文字符，英文字母、数字及“_”、“-”组成，且长度为[1-64]个字符。
	Name *string `json:"name,omitempty"`

	// **参数说明**：设备接入实例的描述信息。 **取值范围**：长度不超过256，只允许中文、字母、数字、以及_，,.。、&-等字符的组合
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
