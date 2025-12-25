package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcEndpointAddress Vpc 端点地址
type VpcEndpointAddress struct {
}

func (o VpcEndpointAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcEndpointAddress struct{}"
	}

	return strings.Join([]string{"VpcEndpointAddress", string(data)}, " ")
}
