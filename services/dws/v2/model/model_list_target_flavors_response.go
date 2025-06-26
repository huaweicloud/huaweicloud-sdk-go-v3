package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTargetFlavorsResponse Response Object
type ListTargetFlavorsResponse struct {

	// **参数解释**： 规格数量。 **取值范围**： 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 规格详情列表。接口返回的规格列表最多为20条。 **取值范围**： 不涉及。
	Flavors        *[]FlavorInfoResponse `json:"flavors,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListTargetFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTargetFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListTargetFlavorsResponse", string(data)}, " ")
}
