package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelAssetsRequest Request Object
type ListModelAssetsRequest struct {

	// **参数解释**： 资产编码。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetCode *string `json:"asset_code,omitempty"`

	// **参数解释**： 资产来源。 **约束限制**： 不涉及 **取值范围**： * Preset：预置 * AIGallery：订阅 * Import：导入 * Publish：发布 **默认取值**： 不涉及
	AssetSource *string `json:"asset_source,omitempty"`

	// **参数解释**： 资产类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetType *string `json:"asset_type,omitempty"`

	// **参数解释**： 资产子类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SubAssetType *string `json:"sub_asset_type,omitempty"`

	// **参数解释**： 对话id。 **约束限制**： 不涉及 **取值范围**： 长度为[1-64]个字符。 **默认取值**： 不涉及
	ChatId *string `json:"chat_id,omitempty"`

	// **参数解释**： 资产应用场景。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetActions *[]string `json:"asset_actions,omitempty"`

	// 模型名称，支持模糊匹配
	AssetName *string `json:"asset_name,omitempty"`

	// **参数解释**： 偏移量。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,100000000]。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 返回限制个数。 **约束限制**： 不涉及 **取值范围**： [1-1000] **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 排序规则。 **约束限制**： 不涉及 **取值范围**： - DESC：降序。 - ASC：升序。 **默认取值**： DESC
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListModelAssetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelAssetsRequest struct{}"
	}

	return strings.Join([]string{"ListModelAssetsRequest", string(data)}, " ")
}
