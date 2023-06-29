package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2xEdgesResponse Response Object
type ListV2xEdgesResponse struct {

	// **参数说明**：总数。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：数据列表。
	Edges          *[]V2XEdgeListResponseDto `json:"edges,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListV2xEdgesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2xEdgesResponse struct{}"
	}

	return strings.Join([]string{"ListV2xEdgesResponse", string(data)}, " ")
}
