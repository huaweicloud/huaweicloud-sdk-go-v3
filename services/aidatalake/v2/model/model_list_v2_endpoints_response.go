package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2EndpointsResponse Response Object
type ListV2EndpointsResponse struct {

	// **参数解释**：总记录数。 **取值范围**：不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：端点列表。
	Endpoints      *[]V2EndpointBriefInfo `json:"endpoints,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListV2EndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2EndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListV2EndpointsResponse", string(data)}, " ")
}
