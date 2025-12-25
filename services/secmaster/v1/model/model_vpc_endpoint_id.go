package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcEndpointId Vpc 端点ID
type VpcEndpointId struct {
}

func (o VpcEndpointId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcEndpointId struct{}"
	}

	return strings.Join([]string{"VpcEndpointId", string(data)}, " ")
}
