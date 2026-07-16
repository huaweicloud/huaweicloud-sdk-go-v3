package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetModel **参数解释**：自定义训练作业产物发布成模型的信息。
type AssetModel struct {

	// **参数解释**：模型名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：模型名称。 **取值范围**：不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**：模型发布版本。 **取值范围**：不涉及。
	Version string `json:"version"`

	// **参数解释**：模型描述。 **取值范围**：不涉及。
	Desc *string `json:"desc,omitempty"`

	// **参数解释**：模型品牌。 **取值范围**：不涉及。
	Series *string `json:"series,omitempty"`

	// **参数解释**：模型类型。 **取值范围**：不涉及。
	Type string `json:"type"`

	// **参数解释**：模型资产描述。\\n**取值范围**：不涉及。
	ModelDesc *string `json:"model_desc,omitempty"`

	// **参数解释**：父资产ID（可选），选择已有模型时传递。\\n**取值范围**：不涉及。
	ParentAssetId *string `json:"parent_asset_id,omitempty"`
}

func (o AssetModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetModel struct{}"
	}

	return strings.Join([]string{"AssetModel", string(data)}, " ")
}
