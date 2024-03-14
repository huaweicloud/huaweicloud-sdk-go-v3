package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointsResponse Response Object
type ListEndpointsResponse struct {

	// 终端节点列表。
	Endpoints *[]EndpointResponseBody `json:"endpoints,omitempty"`

	// 满足查询条件的终端节点总条数，不受分页（即limit、offset参数）影响。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEndpointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointsResponse", string(data)}, " ")
}
