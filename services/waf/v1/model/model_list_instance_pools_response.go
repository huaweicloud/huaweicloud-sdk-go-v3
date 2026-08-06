package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancePoolsResponse Response Object
type ListInstancePoolsResponse struct {

	// **参数解释：** 符合条件的实例组总数。 **取值范围：** 不涉及
	Total *int64 `json:"total,omitempty"`

	// **参数解释：** 实例组列表，详细字段说明请参见PremiumWafPoolResponse object。 **取值范围：** 不涉及
	Items          *[]PremiumWafPoolResponse `json:"items,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListInstancePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancePoolsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancePoolsResponse", string(data)}, " ")
}
