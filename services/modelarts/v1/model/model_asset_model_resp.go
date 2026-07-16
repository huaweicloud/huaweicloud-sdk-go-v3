package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssetModelResp **参数解释**：自定义训练作业产物发布成模型的信息。
type AssetModelResp struct {

	// **参数解释**：模型id。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：模型名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：模型名称。 **取值范围**：不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**：模型发布版本。 **取值范围**：不涉及。
	Version string `json:"version"`

	// **参数解释**：模型发布地址。 **取值范围**：不涉及。
	Location *string `json:"location,omitempty"`

	// **参数解释**：模型描述。 **取值范围**：不涉及。
	Desc *string `json:"desc,omitempty"`

	// **参数解释**：模型品牌。 **取值范围**：不涉及。
	Series *string `json:"series,omitempty"`

	// **参数解释**：模型类型。 **取值范围**：不涉及。
	Type string `json:"type"`
}

func (o AssetModelResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetModelResp struct{}"
	}

	return strings.Join([]string{"AssetModelResp", string(data)}, " ")
}
