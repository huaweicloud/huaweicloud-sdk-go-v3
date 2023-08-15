package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEndpointRequest Request Object
type DeleteEndpointRequest struct {

	// 指定查询访问端点的id
	EndpointId string `json:"endpoint_id"`
}

func (o DeleteEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndpointRequest", string(data)}, " ")
}
