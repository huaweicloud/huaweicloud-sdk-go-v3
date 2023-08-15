package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointRequest Request Object
type UpdateEndpointRequest struct {

	// 指定查询访问端点的id
	EndpointId string `json:"endpoint_id"`

	Body *EndpointUpdateReq `json:"body,omitempty"`
}

func (o UpdateEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRequest", string(data)}, " ")
}
