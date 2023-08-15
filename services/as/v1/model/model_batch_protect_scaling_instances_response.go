package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchProtectScalingInstancesResponse Response Object
type BatchProtectScalingInstancesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchProtectScalingInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchProtectScalingInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchProtectScalingInstancesResponse", string(data)}, " ")
}
