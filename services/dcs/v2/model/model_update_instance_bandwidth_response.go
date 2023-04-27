package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateInstanceBandwidthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceBandwidthResponse", string(data)}, " ")
}
