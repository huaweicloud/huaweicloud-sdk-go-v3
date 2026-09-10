package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelAssetReq 更新模型资产请求体
type UpdateModelAssetReq struct {

	// **参数解释**： 资产描述。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetDesc *string `json:"asset_desc,omitempty"`

	// **参数解释**： 资产名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetName *string `json:"asset_name,omitempty"`
}

func (o UpdateModelAssetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelAssetReq struct{}"
	}

	return strings.Join([]string{"UpdateModelAssetReq", string(data)}, " ")
}
