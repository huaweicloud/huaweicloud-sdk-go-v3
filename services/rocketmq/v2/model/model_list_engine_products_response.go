package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEngineProductsResponse Response Object
type ListEngineProductsResponse struct {

	// **参数解释**： 总数。 **取值范围**： 大于等于0。
	Total float32 `json:"total,omitempty"`

	// **参数解释**： 下个分页的offset。 **取值范围**： 大于等于0。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// **参数解释**： 上个分页的offset。 **取值范围**： 大于等于0。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// **参数解释**： 引擎类型。 **取值范围**： 不涉及。
	Engine *string `json:"engine,omitempty"`

	// **参数解释**： 支持的版本。
	Versions *[]string `json:"versions,omitempty"`

	// **参数解释**： 产品详情列表。
	Products       *[]ProductEntity `json:"products,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListEngineProductsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEngineProductsResponse struct{}"
	}

	return strings.Join([]string{"ListEngineProductsResponse", string(data)}, " ")
}
