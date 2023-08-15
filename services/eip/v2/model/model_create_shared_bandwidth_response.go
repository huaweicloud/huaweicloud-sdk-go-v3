package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSharedBandwidthResponse Response Object
type CreateSharedBandwidthResponse struct {
	Bandwidth      *BandwidthResp `json:"bandwidth,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateSharedBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedBandwidthResponse struct{}"
	}

	return strings.Join([]string{"CreateSharedBandwidthResponse", string(data)}, " ")
}
