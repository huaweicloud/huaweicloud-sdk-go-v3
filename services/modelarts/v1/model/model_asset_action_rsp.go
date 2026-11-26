package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetActionRsp **参数解释**：资产可执行操作详情配置。
type AssetActionRsp struct {

	// **参数解释**：资产action类型。 **取值范围**：不涉及。
	ActionType *string `json:"action_type,omitempty"`

	// **参数解释**：资产action配置文件地址。 **取值范围**：不涉及。
	ActionDir *string `json:"action_dir,omitempty"`
}

func (o AssetActionRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetActionRsp struct{}"
	}

	return strings.Join([]string{"AssetActionRsp", string(data)}, " ")
}
