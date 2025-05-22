package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDssPoolsResponse Response Object
type ListDssPoolsResponse struct {

	// **参数解释**： 专属分布式存储池详情列表。 **取值范围**： 不涉及。
	Pools *[]DssPool `json:"pools,omitempty"`

	// **参数解释**： 专属分布式存储池个数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDssPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDssPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListDssPoolsResponse", string(data)}, " ")
}
