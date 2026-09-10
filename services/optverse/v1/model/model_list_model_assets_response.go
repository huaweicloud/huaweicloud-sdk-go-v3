package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelAssetsResponse Response Object
type ListModelAssetsResponse struct {

	// **参数解释**： 资产列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Assets *[]ModelAssetRsp `json:"assets,omitempty"`

	// **参数解释**： 资产总数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListModelAssetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelAssetsResponse struct{}"
	}

	return strings.Join([]string{"ListModelAssetsResponse", string(data)}, " ")
}
