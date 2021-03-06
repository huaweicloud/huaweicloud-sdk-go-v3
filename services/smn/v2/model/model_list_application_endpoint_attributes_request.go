package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApplicationEndpointAttributesRequest struct {
	// Endpoint的唯一资源标识，可通过[查询Application的Endpoint列表](https://support.huaweicloud.com/api-smn/smn_api_58004.html)获取该标识。

	EndpointUrn string `json:"endpoint_urn"`
}

func (o ListApplicationEndpointAttributesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointAttributesRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointAttributesRequest", string(data)}, " ")
}
