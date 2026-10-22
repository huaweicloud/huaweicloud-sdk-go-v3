package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2AllEndpointsResponse Response Object
type ListV2AllEndpointsResponse struct {

	// **参数解释**：总记录数。 **取值范围**：不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：端点列表。
	Endpoints      *[]V2EndpointSpaceInfo `json:"endpoints,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListV2AllEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2AllEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListV2AllEndpointsResponse", string(data)}, " ")
}
