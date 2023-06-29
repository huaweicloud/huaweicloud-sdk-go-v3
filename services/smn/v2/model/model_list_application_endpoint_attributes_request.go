package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationEndpointAttributesRequest Request Object
type ListApplicationEndpointAttributesRequest struct {

	// Endpoint的唯一资源标识，可通过[查询Application的Endpoint列表](smn_api_58004.xml)获取该标识。
	EndpointUrn string `json:"endpoint_urn"`
}

func (o ListApplicationEndpointAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointAttributesRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointAttributesRequest", string(data)}, " ")
}
