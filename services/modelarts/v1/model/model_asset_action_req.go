package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetActionReq **参数解释**：资产可执行操作详情配置。 **约束限制**：不涉及。
type AssetActionReq struct {

	// **参数解释**：资产action类型。 **约束限制**：0-32位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ActionType *string `json:"action_type,omitempty"`

	// **参数解释**：资产action配置文件地址。 **约束限制**：0-32位字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ActionDir *string `json:"action_dir,omitempty"`
}

func (o AssetActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetActionReq struct{}"
	}

	return strings.Join([]string{"AssetActionReq", string(data)}, " ")
}
