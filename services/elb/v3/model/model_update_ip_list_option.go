package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpListOption **参数解释**：更新IP地址组IP列表的请求参数。  **约束限制**：不涉及
type UpdateIpListOption struct {

	// **参数解释**：IP地址组的名称。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：IP地址组中包含的IP列表。 只支持添加新的IP地址到IP地址组的IP列表中，或更新已有IP地址的描述。不会删除或修改ip_list中已有的IP地址。  **约束限制**：不涉及
	IpList *[]UpdateIpGroupIpOption `json:"ip_list,omitempty"`

	// **参数解释**：备注信息。  **约束限制**：不涉及  **取值范围**：长度为0-255个字符。  **默认取值**：不涉及
	Description *string `json:"description,omitempty"`
}

func (o UpdateIpListOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListOption struct{}"
	}

	return strings.Join([]string{"UpdateIpListOption", string(data)}, " ")
}
