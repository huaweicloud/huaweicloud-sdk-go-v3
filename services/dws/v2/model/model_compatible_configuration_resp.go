package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompatibleConfigurationResp **参数解释**： 参数组信息。 **取值范围**： 不涉及。
type CompatibleConfigurationResp struct {

	// **参数解释**： 参数组ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 参数组名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 连接信息。 **取值范围**： 不涉及。
	Links *[]LinkResp `json:"links,omitempty"`
}

func (o CompatibleConfigurationResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompatibleConfigurationResp struct{}"
	}

	return strings.Join([]string{"CompatibleConfigurationResp", string(data)}, " ")
}
