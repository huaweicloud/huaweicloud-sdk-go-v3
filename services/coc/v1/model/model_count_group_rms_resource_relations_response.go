package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountGroupRmsResourceRelationsResponse Response Object
type CountGroupRmsResourceRelationsResponse struct {

	// **参数解释：** 云资源数量。 **取值范围：** 不涉及。
	Data           *int64 `json:"data,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountGroupRmsResourceRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGroupRmsResourceRelationsResponse struct{}"
	}

	return strings.Join([]string{"CountGroupRmsResourceRelationsResponse", string(data)}, " ")
}
