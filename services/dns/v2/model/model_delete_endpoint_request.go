package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEndpointRequest struct {

	// 待删除endpoint的ID。
	EndpointId string `json:"endpoint_id"`
}

func (o DeleteEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndpointRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndpointRequest", string(data)}, " ")
}
