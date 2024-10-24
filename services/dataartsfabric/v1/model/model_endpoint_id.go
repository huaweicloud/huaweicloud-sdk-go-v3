package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointId Endpoint Id，32~36位的英文、数字、短横组合。
type EndpointId struct {
}

func (o EndpointId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointId struct{}"
	}

	return strings.Join([]string{"EndpointId", string(data)}, " ")
}
