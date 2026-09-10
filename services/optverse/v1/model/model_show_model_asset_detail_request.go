package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModelAssetDetailRequest Request Object
type ShowModelAssetDetailRequest struct {

	// **参数解释**： 资产ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetId string `json:"asset_id"`
}

func (o ShowModelAssetDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModelAssetDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowModelAssetDetailRequest", string(data)}, " ")
}
