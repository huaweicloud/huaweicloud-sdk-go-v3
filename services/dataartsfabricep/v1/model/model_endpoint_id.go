package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointId endpoint空间id
type EndpointId struct {
}

func (o EndpointId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointId struct{}"
	}

	return strings.Join([]string{"EndpointId", string(data)}, " ")
}
