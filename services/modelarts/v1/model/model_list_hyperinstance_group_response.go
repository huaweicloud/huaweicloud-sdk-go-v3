package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHyperinstanceGroupResponse Response Object
type ListHyperinstanceGroupResponse struct {

	// **参数解释**：当前页数。 **取值范围**：1 - 2097152
	Current *int32 `json:"current,omitempty"`

	// **参数解释**：总的页数。 **取值范围**：1 - 2097152
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释**：每一页的数量。设置查询时每页返回的最大实例数量，此值将直接影响到返回数据的量。调整此值可以优化查询性能或提高单次请求的信息量。 **取值范围**：1 - 1024
	Size *int32 `json:"size,omitempty"`

	// **参数解释**：总的记录数量。表示当前查询条件下，满足条件的实例总数。此字段用于用户了解数据的整体规模。 **取值范围**：1 - 2147483647
	Total *int64 `json:"total,omitempty"`

	// **参数解释**：hyperinstance-group详情列表。 **默认取值**：不涉及。
	Data           *[]HyperinstanceGroupResponse `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListHyperinstanceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHyperinstanceGroupResponse struct{}"
	}

	return strings.Join([]string{"ListHyperinstanceGroupResponse", string(data)}, " ")
}
