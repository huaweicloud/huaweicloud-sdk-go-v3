package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetResponse Response Object
type CreateAssetResponse struct {

	// **参数解释**：资产id。 **取值范围**：不涉及。
	AssetId        *string `json:"asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetResponse", string(data)}, " ")
}
