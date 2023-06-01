package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteBandwidthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DeleteBandwidthResponse", string(data)}, " ")
}
