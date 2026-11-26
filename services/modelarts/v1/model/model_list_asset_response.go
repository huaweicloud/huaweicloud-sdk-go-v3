package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetResponse Response Object
type ListAssetResponse struct {

	// **参数解释**：符合条件的资产总数量。 **取值范围**：不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：树节点列表及关系。
	Data           *[]AssetDetail `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetResponse struct{}"
	}

	return strings.Join([]string{"ListAssetResponse", string(data)}, " ")
}
