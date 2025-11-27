package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourcesOfResourceViewResponse Response Object
type CountResourcesOfResourceViewResponse struct {

	// **参数解释：** 云资源数量。 **取值范围：** 不涉及。
	Data           *int64 `json:"data,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountResourcesOfResourceViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesOfResourceViewResponse struct{}"
	}

	return strings.Join([]string{"CountResourcesOfResourceViewResponse", string(data)}, " ")
}
