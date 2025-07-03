package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlowsResponse Response Object
type ListFlowsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 流列表
	Flows          *[]ListFlowRespItem `json:"flows,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListFlowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowsResponse struct{}"
	}

	return strings.Join([]string{"ListFlowsResponse", string(data)}, " ")
}
