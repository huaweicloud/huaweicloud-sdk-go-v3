package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBandwidthResponse struct {
	Bandwidth      *Bandwidth `json:"bandwidth,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthResponse struct{}"
	}

	return strings.Join([]string{"ShowBandwidthResponse", string(data)}, " ")
}
