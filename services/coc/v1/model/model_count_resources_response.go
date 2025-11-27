package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourcesResponse Response Object
type CountResourcesResponse struct {

	// **参数解释：** 云资源数量。 **取值范围：** 不涉及。
	Data *int64 `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesResponse struct{}"
	}

	return strings.Join([]string{"CountResourcesResponse", string(data)}, " ")
}
