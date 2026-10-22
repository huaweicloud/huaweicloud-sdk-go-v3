package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2TagResourcesResponse Response Object
type ListV2TagResourcesResponse struct {

	// **参数解释**：资源实例列表。
	Resources *[]V2TagResource `json:"resources,omitempty"`

	// **参数解释**：资源实例总数。 **取值范围**：0~2147483647。
	TotalCount *int32 `json:"total_count,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListV2TagResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2TagResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListV2TagResourcesResponse", string(data)}, " ")
}
