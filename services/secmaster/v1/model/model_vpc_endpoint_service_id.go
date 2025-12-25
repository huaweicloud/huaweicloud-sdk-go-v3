package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcEndpointServiceId Vpc端点服务ID
type VpcEndpointServiceId struct {
}

func (o VpcEndpointServiceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcEndpointServiceId struct{}"
	}

	return strings.Join([]string{"VpcEndpointServiceId", string(data)}, " ")
}
