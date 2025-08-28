package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupOption **参数解释**：更新IP地址组请求参数。  **约束限制**：不涉及
type UpdateIpGroupOption struct {

	// **参数解释**：IP地址组的描述信息。  **约束限制**：不涉及  **取值范围**：长度为0-255个字符。  **默认取值**：不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：IP地址组的名称  **约束限制**：不涉及  **取值范围**：长度为0-255个字符。  **默认取值**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：IP地址组的IP列表。 注意：只支持全量更新，即IP地址组中的ip_list将被全量覆盖，不在请求参数中的IP地址将被移除。  **约束限制**：不涉及
	IpList *[]UpdateIpGroupIpOption `json:"ip_list,omitempty"`
}

func (o UpdateIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupOption", string(data)}, " ")
}
