package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteSharedBandwidthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSharedBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSharedBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DeleteSharedBandwidthResponse", string(data)}, " ")
}
