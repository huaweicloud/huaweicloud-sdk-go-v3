package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListEndpointServiceResponse struct {
	// 终端节点服务列表

	EndpointServices *[]ServiceDetailsResp `json:"endpoint_services,omitempty"`
	// 满足查询条件的终端节点服务总条数，不受分页（即limit、offset参数）影响。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEndpointServiceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointServiceResponse", string(data)}, " ")
}
