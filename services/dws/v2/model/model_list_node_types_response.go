package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodeTypesResponse Response Object
type ListNodeTypesResponse struct {

	// **参数解释**： 可用的规格列表。 **取值范围**： 非空对象列表。
	NodeTypes *[]NodeTypes `json:"node_types,omitempty"`

	// **参数解释**： 规格总数 **取值范围**： 大于等于0的正整数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNodeTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeTypesResponse struct{}"
	}

	return strings.Join([]string{"ListNodeTypesResponse", string(data)}, " ")
}
