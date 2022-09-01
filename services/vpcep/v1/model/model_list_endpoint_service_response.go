package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEndpointServiceResponse struct {

	// 终端节点服务列表
	EndpointServices *[]ServiceList `json:"endpoint_services,omitempty" xml:"endpoint_services"`

	// 满足查询条件的终端节点服务总条数，不受分页（即limit、offset参数）影响。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointServiceResponse", string(data)}, " ")
}
