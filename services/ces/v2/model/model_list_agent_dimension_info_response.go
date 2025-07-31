package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentDimensionInfoResponse Response Object
type ListAgentDimensionInfoResponse struct {

	// **参数解释**： 维度信息。
	Dimensions *[]AgentDimension `json:"dimensions,omitempty"`

	// **参数解释**： 维度信息总数。 **取值范围**： 整数，最小值为0，最大值为2147483647。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAgentDimensionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentDimensionInfoResponse struct{}"
	}

	return strings.Join([]string{"ListAgentDimensionInfoResponse", string(data)}, " ")
}
