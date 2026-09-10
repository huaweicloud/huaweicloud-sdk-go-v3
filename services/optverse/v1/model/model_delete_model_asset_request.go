package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteModelAssetRequest Request Object
type DeleteModelAssetRequest struct {

	// **参数解释**： 资产ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssetId string `json:"asset_id"`
}

func (o DeleteModelAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteModelAssetRequest struct{}"
	}

	return strings.Join([]string{"DeleteModelAssetRequest", string(data)}, " ")
}
