package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAgentDimensionInfoResponse struct {

	// 维度信息
	Dimensions *[]AgentDimension `json:"dimensions,omitempty" xml:"dimensions"`

	// 维度信息总数
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAgentDimensionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentDimensionInfoResponse struct{}"
	}

	return strings.Join([]string{"ListAgentDimensionInfoResponse", string(data)}, " ")
}
