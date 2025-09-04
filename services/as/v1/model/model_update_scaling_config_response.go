package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScalingConfigResponse Response Object
type UpdateScalingConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateScalingConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateScalingConfigResponse", string(data)}, " ")
}
