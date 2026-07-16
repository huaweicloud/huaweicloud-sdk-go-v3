package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDynamicStoragesResponse Response Object
type ListDynamicStoragesResponse struct {

	// **参数解释**：当前页数。 **取值范围**：正整数。
	Current *int32 `json:"current,omitempty"`

	// **参数解释**：动态挂载存储实例数据。
	Data *[]DataVolumesRes `json:"data,omitempty"`

	// **参数解释**：总的页数。 **取值范围**：正整数。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释**：每一页的数量。 **取值范围**：正整数。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**：总的记录数量。 **取值范围**：非负整数。
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDynamicStoragesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDynamicStoragesResponse struct{}"
	}

	return strings.Join([]string{"ListDynamicStoragesResponse", string(data)}, " ")
}
