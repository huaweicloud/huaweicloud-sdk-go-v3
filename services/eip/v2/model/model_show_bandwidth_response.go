package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBandwidthResponse Response Object
type ShowBandwidthResponse struct {
	Bandwidth      *BandwidthResp `json:"bandwidth,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthResponse struct{}"
	}

	return strings.Join([]string{"ShowBandwidthResponse", string(data)}, " ")
}
