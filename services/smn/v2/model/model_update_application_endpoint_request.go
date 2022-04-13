package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateApplicationEndpointRequest struct {
	// Endpoint的唯一资源标识，可通过[查询Application的Endpoint列表](https://support.huaweicloud.com/api-smn/ListApplicationEndpoints.html)获取该标识。

	EndpointUrn string `json:"endpoint_urn"`

	Body *UpdateApplicationEndpointRequestBody `json:"body,omitempty"`
}

func (o UpdateApplicationEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationEndpointRequest struct{}"
	}

	return strings.Join([]string{"UpdateApplicationEndpointRequest", string(data)}, " ")
}
