package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetResponse Response Object
type DeleteAssetResponse struct {

	// **参数解释**：资产id。 **取值范围**：不涉及。
	AssetId        *string `json:"asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetResponse", string(data)}, " ")
}
