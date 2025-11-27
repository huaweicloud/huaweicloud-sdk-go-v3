package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceViewsResponse Response Object
type UpdateResourceViewsResponse struct {

	// **参数解释：** 视图id。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateResourceViewsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceViewsResponse struct{}"
	}

	return strings.Join([]string{"UpdateResourceViewsResponse", string(data)}, " ")
}
