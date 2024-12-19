package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBandwidthResponse Response Object
type UpdateBandwidthResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthResponse", string(data)}, " ")
}
