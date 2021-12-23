package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateBandwidthResponse struct {
	Bandwidth      *BandwidthResp `json:"bandwidth,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthResponse", string(data)}, " ")
}
