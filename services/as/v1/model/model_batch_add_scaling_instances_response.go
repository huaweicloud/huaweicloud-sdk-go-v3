package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAddScalingInstancesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddScalingInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddScalingInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchAddScalingInstancesResponse", string(data)}, " ")
}
