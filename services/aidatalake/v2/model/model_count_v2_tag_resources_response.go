package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountV2TagResourcesResponse Response Object
type CountV2TagResourcesResponse struct {

	// **参数解释**：标签资源实例数量。 **取值范围**：0~2147483647。
	TotalCount *int32 `json:"total_count,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountV2TagResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountV2TagResourcesResponse struct{}"
	}

	return strings.Join([]string{"CountV2TagResourcesResponse", string(data)}, " ")
}
