package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListServicePublicDetailsResponse struct {

	// 终端节点服务列表。
	EndpointServices *[]EndpointService `json:"endpoint_services,omitempty"`

	// 满足查询条件的公共终端节点服务总条数，不受分页（即limit、offset参数）影响。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListServicePublicDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicePublicDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServicePublicDetailsResponse", string(data)}, " ")
}
