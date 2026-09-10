package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachSubNetworkInterfaceOption struct {

	// **参数解释**： 辅助弹性网卡所挂载的弹性网卡的ID。 **取值范围**： 带“-”的标准UUID格式。
	ParentId string `json:"parent_id"`

	// **参数解释**： 辅助弹性网卡关联的安全组的ID列表。例如：\"security_groups\": [\"a0608cbf-d047-4f54-8b28-cd7b59853fff\"]。 **取值范围**： 如果请求时不指定此参数，辅助弹性网卡创建后会自动关联默认安全组。
	SecurityGroups *[]string `json:"security_groups,omitempty"`

	// **参数解释**： 辅助弹性网卡安全使能标记，如果不使能则安全组不生效。 **取值范围**： 不涉及。
	SecurityEnabled *bool `json:"security_enabled,omitempty"`
}

func (o AttachSubNetworkInterfaceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachSubNetworkInterfaceOption struct{}"
	}

	return strings.Join([]string{"AttachSubNetworkInterfaceOption", string(data)}, " ")
}
